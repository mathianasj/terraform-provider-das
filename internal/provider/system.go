package provider

import (
	"bytes"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	systems "github.com/mathianasj/terraform-provider-das/client/systems"
	"github.com/mathianasj/terraform-provider-das/internal/util"
	models "github.com/mathianasj/terraform-provider-das/models"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func decisionMappingColumn() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "column key (also the search key)",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dot-separated decision property path",
			},
			"type": {
				Type:        schema.TypeString,
				Description: "column type: one of \"string\", \"boolean\", \"date\", \"integer\", \"float\"",
				Optional:    true,
			},
		},
	}
}

func HashDecisionMappingAllowed(m map[string]interface{}, buf bytes.Buffer) {
	if v, ok := m["path"]; ok {
		buf.WriteString(fmt.Sprintf("%v", v.(string)))
	}

	if v, ok := m["exected"]; ok {
		buf.WriteString(fmt.Sprintf("%v", v.(string)))
	}

	if v, ok := m["negated"]; ok {
		buf.WriteString(fmt.Sprintf("%v", v.(bool)))
	}
}

func decisionMappingAllowed() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dot-separated decision property path",
			},
			"expected": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"negated": {
				Type:        schema.TypeBool,
				Description: "when set to true, decision is Allowed when the mapped property IS NOT equal to the expected value",
				Optional:    true,
			},
		},
	}
}

func decisionMappingReason() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dot-separated decision property path",
			},
		},
	}
}

func decisionMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allowed": {
				Type:     schema.TypeSet,
				Elem:     decisionMappingAllowed(),
				Optional: true,
				MaxItems: 1,
				Set:      schema.HashResource(decisionMappingAllowed()),
			},
			"columns": {
				Type:     schema.TypeSet,
				Elem:     decisionMappingColumn(),
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeSet,
				Elem:     decisionMappingReason(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceSystem() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "system name",
				ValidateFunc: validateName,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "system type e.g. kubernetes",
			},
			"decision_mapping": {
				Type:        schema.TypeSet,
				Description: "location of key attributes and additional columns in the decisions grouped by policy entry point path",
				Elem:        decisionMapping(),
				Optional:    true,
				Set:         schema.HashResource(decisionMapping()),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	log.Print("[INFO] input decions : ", d.Get("decision_mapping"))
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	options := systems.NewCreateSystemParams().WithBody(&models.SystemsV1SystemsPostRequest{
		Name:             util.GetStringRef(d, "name"),
		Type:             util.GetStringRef(d, "type"),
		Description:      util.GetString(d, "description"),
		DecisionMappings: mapDecisionMapping(util.GetSet(d, "decision_mapping")),
	})

	resp, err := apiClient.Systems.CreateSystem(options, bearerTokenAuth)

	if err != nil {
		return err
	}

	mapResult(d, resp.Payload.Result)

	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] provider: testing")
	log.Println("[INFO] das: getting current system ", d.Id())
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	systemId := d.Id()
	getSystemParams := systems.NewGetSystemParams().WithSystem(systemId)

	resp, err := apiClient.Systems.GetSystem(getSystemParams, bearerTokenAuth)

	if err != nil {
		return fmt.Errorf("error finding Item with ID %s", systemId)
	}

	mapResult(d, resp.Payload.Result)

	return nil
}

func mapResult(d *schema.ResourceData, result *models.SystemsV1SystemConfig) error {
	d.SetId(*result.ID)
	d.Set("name", result.Name)
	d.Set("type", result.Type)
	d.Set("description", result.Description)

	log.Printf("[INFO] before set : ", d.Get("decision_mapping").(*schema.Set))
	err := d.Set("decision_mapping", convertDecisionMapping(result.DecisionMappings))
	if err != nil {
		log.Printf("[INFO] error setting decision : ", err)
	}
	log.Printf("[INFO] after set : ", d.Get("decision_mapping"))

	return nil
}

func convertDecisionMapping(systemsV1RuleDecisionMappings map[string]models.SystemsV1RuleDecisionMappings) []interface{} {
	var decisionMappings []interface{}

	log.Println("[INFO] mapping decision mappings")

	for i, o := range systemsV1RuleDecisionMappings {
		log.Println("[INFO] index : ", i)
		decisionMapping := map[string]interface{}{
			"name":    i,
			"allowed": convertAllowed(o.Allowed),
			"reason":  convertReason(o.Reason),
			"columns": convertColumns(o.Columns),
		}
		log.Println("[INFO] output mapping : ", decisionMapping)

		decisionMappings = append(decisionMappings, decisionMapping)
	}

	log.Println("[INFO] output mappings : ", decisionMappings)
	return decisionMappings
}

func mapAllowed(allowed *schema.Set) *models.SystemsV1AllowedMapping {
	typed := allowed.List()[0].(map[string]interface{})
	var expected interface{}

	if typed["expected"].(string) == "true" {
		expected = true
	} else if typed["expected"].(string) == "false" {
		expected = false
	} else {
		expected = typed["expected"].(string)
	}

	return &models.SystemsV1AllowedMapping{
		Expected: expected,
		Path:     util.ConvertStringRef(typed["path"].(string)),
		Negated:  util.ConvertBoolRef(typed["negated"].(bool)),
	}
}

func mapReason(reason *schema.Set) *models.SystemsV1ReasonMapping {
	if len(reason.List()) > 0 {
		typed := reason.List()[0].(map[string]interface{})

		return &models.SystemsV1ReasonMapping{
			Path: util.ConvertStringRef(typed["path"].(string)),
		}
	}
	return nil
}

func mapColumns(columns *schema.Set) []*models.SystemsV1ColumnMapping {
	var out []*models.SystemsV1ColumnMapping

	for _, o := range columns.List() {
		typed := o.(map[string]interface{})

		out = append(out, &models.SystemsV1ColumnMapping{
			Key:  util.ConvertStringRef(typed["key"].(string)),
			Path: util.ConvertStringRef(typed["path"].(string)),
			Type: util.ConvertStringRef(typed["type"].(string)),
		})
	}

	return out
}

func mapDecisionMapping(decisions *schema.Set) map[string]models.SystemsV1RuleDecisionMappings {
	out := map[string]models.SystemsV1RuleDecisionMappings{}

	for _, o := range decisions.List() {
		typed := o.(map[string]interface{})

		out[typed["name"].(string)] = models.SystemsV1RuleDecisionMappings{
			Allowed: mapAllowed(typed["allowed"].(*schema.Set)),
			Columns: mapColumns(typed["columns"].(*schema.Set)),
			Reason:  mapReason(typed["reason"].(*schema.Set)),
		}
	}

	return out
}

func convertAllowed(systemsV1AllowedMapping *models.SystemsV1AllowedMapping) []map[string]interface{} {
	allowedMappings := []map[string]interface{}{
		{
			"path":     *systemsV1AllowedMapping.Path,
			"negated":  *systemsV1AllowedMapping.Negated,
			"expected": fmt.Sprintf("%v", systemsV1AllowedMapping.Expected),
		},
	}

	return allowedMappings
}

func convertReason(systemsV1ReasonMapping *models.SystemsV1ReasonMapping) []map[string]interface{} {
	var reasons []map[string]interface{}

	if systemsV1ReasonMapping == nil {
		reasons = []map[string]interface{}{}
	} else {
		reasons = []map[string]interface{}{
			{
				"path": *systemsV1ReasonMapping.Path,
			},
		}
	}

	return reasons
}

func convertColumns(systemsV1ColumnMapping []*models.SystemsV1ColumnMapping) []map[string]interface{} {
	var columns []map[string]interface{}

	for _, o := range systemsV1ColumnMapping {
		columns = append(columns, map[string]interface{}{
			"key":  *o.Key,
			"path": *o.Path,
			"type": *o.Type,
		})
	}

	return columns
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	systemId := d.Id()

	updateSystemParams := systems.NewUpdateSystemParams().WithBody(&models.SystemsV1SystemsPutRequest{
		Name:             util.GetStringRef(d, "name"),
		Type:             util.GetStringRef(d, "type"),
		Description:      util.GetString(d, "description"),
		DecisionMappings: mapDecisionMapping(util.GetSet(d, "decision_mapping")),
	}).WithSystem(systemId)

	_, err := apiClient.Systems.UpdateSystem(updateSystemParams, bearerTokenAuth)

	if err != nil {
		return err
	}

	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	systemId := d.Id()

	deleteSystemParams := systems.NewDeleteSystemParams().WithSystem(systemId)

	_, err := apiClient.Systems.DeleteSystem(deleteSystemParams, bearerTokenAuth)

	if err != nil {
		return err
	}
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	systemId := d.Id()
	getSystemParams := systems.NewGetSystemParams().WithSystem(systemId)

	_, err := apiClient.Systems.GetSystem(getSystemParams, bearerTokenAuth)

	if err != nil {
		return false, err
	}
	return true, nil
}
