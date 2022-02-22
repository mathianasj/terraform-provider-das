package provider

import (
	"log"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	httptransport "github.com/go-openapi/runtime/client"
	apiClient "github.com/mathianasj/terraform-provider-das/client"
)

type TfConfig struct {
	client          *apiClient.StyraAPI
	bearerTokenAuth runtime.ClientAuthInfoWriter
}

func ConvertTfConfig(m interface{}) *TfConfig {
	tfConfig := m.(TfConfig)

	return &tfConfig
}

func (t *TfConfig) GetClientAuth() (*apiClient.StyraAPI, runtime.ClientAuthInfoWriter) {
	return t.client, t.bearerTokenAuth
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"das_system": resourceSystem(),
			"das_policy": resourcePolicy(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Printf("[INFO] loading api client")
	address := d.Get("address").(string)
	token := d.Get("token").(string)

	// create the transport
	transport := httptransport.New(address, "", []string{"https"})
	transport.Producers["*/*"] = runtime.JSONProducer()
	bearerTokenAuth := httptransport.BearerToken(token)

	// create the API client, with the transport
	return TfConfig{
		client:          apiClient.New(transport, strfmt.Default),
		bearerTokenAuth: bearerTokenAuth,
	}, nil
}
