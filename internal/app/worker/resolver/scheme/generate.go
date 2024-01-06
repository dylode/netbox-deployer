package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"dylaan.nl/netbox-deployer/internal/app/worker"
	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/Khan/genqlient/graphql"
)

const defaultConfigFilePath = "config.yaml"

func generateModelsToTypes(client graphql.Client) {
	allTypes, err := netbox.GetAllTypes(context.Background(), client)
	if err != nil {
		panic(err)
	}

	fmt.Println(allTypes)
}

func main() {
	configFilePath := os.Getenv("NBDEPLOY_CONFIG")
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}

	config, err := worker.NewConfigFromPath(configFilePath)
	if err != nil {
		panic(err)
	}

	httpClient := http.Client{}
	graphqlClient := graphql.NewClient(config.Worker.GraphqlURL, &httpClient)

	generateModelsToTypes(graphqlClient)
}
