package provider

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mathianasj/terraform-provider-das/client/policies"
	"github.com/mathianasj/terraform-provider-das/internal/util"
	models "github.com/mathianasj/terraform-provider-das/models"
)

func resourcePolicy() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "system name",
				ValidateFunc: validateName,
			},
			"policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "system name",
			},
			"system": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "system name",
				ValidateFunc: validateName,
			},
			"file": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "system name",
				ValidateFunc: validateName,
			},
		},
		Create: resourceCreatePolicy,
		Read:   resourceReadPolicy,
		Update: resourceUpdatePolicy,
		Delete: resourceDeletePolicy,
		Exists: resourceExistsPolicy,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceUpdatePolicy(d *schema.ResourceData, m interface{}) error {
	return resourceCreatePolicy(d, m)
}

func resourceDeletePolicy(d *schema.ResourceData, m interface{}) error {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	var modules = map[string]*string{}
	modules[d.Get("file").(string)] = nil

	options := policies.NewUpdatePolicyParams().WithPolicy("systems/" + d.Get("system").(string) + "/" + d.Get("path").(string)).WithBody(&models.PoliciesV1PoliciesPutRequest{
		Modules: modules,
	})

	_, err := apiClient.Policies.UpdatePolicy(options, bearerTokenAuth)

	if err != nil {
		typedError := err.(*runtime.APIError)
		log.Printf("[ERROR] error from API ", typedError.Response.(runtime.ClientResponse).Body())
		return err
	}

	d.SetId("/systems/" + d.Get("system").(string) + "/" + d.Get("path").(string) + "/" + d.Get("file").(string))

	return nil
}

func resourceExistsPolicy(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	options := policies.NewGetPolicyParams().WithPolicy("systems/" + d.Get("system").(string) + "/" + d.Get("path").(string))

	log.Printf("[INFO] request policy with ", options)

	_, err := apiClient.Policies.GetPolicy(options, bearerTokenAuth)

	if err != nil {
		typedError := err.(*runtime.APIError)
		log.Printf("[ERROR] error from API ", typedError.Response)
		return false, err
	}

	return true, nil
}

func resourceReadPolicy(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Fetching policy")
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	options := policies.NewGetPolicyParams().WithPolicy("systems/" + d.Get("system").(string) + "/" + d.Get("path").(string))

	resp, err := apiClient.Policies.GetPolicy(options, bearerTokenAuth)

	if err != nil {
		typedError := err.(*runtime.APIError)
		log.Printf("[ERROR] error from API ", typedError.Response)
		return err
	}

	log.Printf("[INFO] read policy ", resp.Payload.Result)
	// typed response body
	typedResponse := resp.Payload.Result.(map[string]interface{})
	modules := typedResponse["modules"].(map[string]interface{})

	d.Set("policy", modules[d.Get("file").(string)])

	return nil
}

func resourceCreatePolicy(d *schema.ResourceData, m interface{}) error {
	apiClient, bearerTokenAuth := ConvertTfConfig(m).GetClientAuth()

	var modules = map[string]*string{}
	modules[d.Get("file").(string)] = util.GetStringRef(d, "policy")

	options := policies.NewUpdatePolicyParams().WithPolicy("systems/" + d.Get("system").(string) + "/" + d.Get("path").(string)).WithBody(&models.PoliciesV1PoliciesPutRequest{
		Modules: modules,
	})

	log.Println("[INFO] request options : ", options)
	log.Println("[INFO] request modules : ", modules)

	resp, err := apiClient.Policies.UpdatePolicy(options, bearerTokenAuth)

	if err != nil {
		typedError := err.(*runtime.APIError)
		log.Printf("[ERROR] error from API ", typedError.Response.(runtime.ClientResponse).Body())
		// log.Println("[ERROR] unable to update request", err)
		return err
	}

	log.Println("[INFO] request id : ", resp.Payload.RequestID)

	d.SetId("/systems/" + d.Get("system").(string) + "/" + d.Get("path").(string) + "/" + d.Get("file").(string))

	return nil
}
