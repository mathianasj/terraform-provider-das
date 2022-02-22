package main

import (
	"log"

	"github.com/hashicorp/terraform/plugin"
	"github.com/mathianasj/terraform-provider-das/internal/provider"
)

func main() {
	logFlags := log.Flags()
	logFlags = logFlags &^ (log.Ldate | log.Ltime)
	log.SetFlags(logFlags)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})

	// 	// create the transport
	// 	transport := httptransport.New("joshua.styra.com", "", []string{"https"})
	// 	transport.Producers["*/*"] = runtime.JSONProducer()
	// 	bearerTokenAuth := httptransport.BearerToken(os.Getenv("API_ACCESS_TOKEN"))

	// 	// create the API client, with the transport
	// 	client := apiClient.New(transport, strfmt.Default)

	// 	// make the request to get all systems
	// 	resp, err := client.Systems.ListSystems(systems.NewListSystemsParams(), bearerTokenAuth)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%#v\n", resp.Payload)
}
