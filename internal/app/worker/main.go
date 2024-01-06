package worker

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dylaan.nl/netbox-deployer/internal/app/worker/api"
	"dylaan.nl/netbox-deployer/internal/app/worker/resolver"
	"github.com/Khan/genqlient/graphql"
)

const errorChanSize = 10

func Run(config Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	// setup
	httpClient := http.Client{}
	graphqlClient := graphql.NewClient(config.Worker.GraphqlURL, &httpClient)

	resolver := resolver.New(
		resolver.NewConfig().WithClient(graphqlClient),
	)
	api := api.New(
		api.NewConfig().
			WithHost(config.Worker.Host).
			WithPort(config.Worker.Port),
		resolver.GetUpdateChan(),
	)

	// run
	errc := make(chan error, errorChanSize)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- api.Run()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- resolver.Run()
	}()

	// shutdown
	var err error
	select {
	case <-ctx.Done():
		fmt.Println("closing due to interrupt")
	case err = <-errc:
		fmt.Println("closing due to error")
	}
	cancel()
	_ = api.Close()
	_ = resolver.Close()
	wg.Wait()

	return err
}
