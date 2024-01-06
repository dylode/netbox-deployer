package resolver

import (
	"fmt"
	"net/http"

	"dylaan.nl/netbox-deployer/internal/pkg/worker"
	"github.com/Khan/genqlient/graphql"
)

const workerUpdateChanBuffer = 1_000

type resolver struct {
	config Config

	graphqlClient *graphql.Client
	updateChan    chan worker.Update
}

func New(config Config) *resolver {
	httpClient := http.Client{}

	graphqlClient := graphql.NewClient(config.graphqlURL, &httpClient)

	return &resolver{
		config: config,

		graphqlClient: &graphqlClient,
		updateChan:    make(chan worker.Update, workerUpdateChanBuffer),
	}
}

func (r resolver) GetUpdateChan() chan worker.Update {
	return r.updateChan
}

func (r resolver) Run() error {
	for update := range r.updateChan {
		fmt.Println(update)
	}

	return nil
}

func (r resolver) Close() error {
	close(r.updateChan)
	return nil
}
