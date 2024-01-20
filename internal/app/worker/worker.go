package worker

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dylaan.nl/netbox-deployer/internal/app/worker/api"
	"dylaan.nl/netbox-deployer/internal/app/worker/state"
	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/Khan/genqlient/graphql"
	"github.com/luthermonson/go-proxmox"
)

const defaultChanSize = 10

func Run(config Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	// setup
	httpClient := http.Client{
		Transport: &authedTransport{
			token:   config.Netbox.Token,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient(config.Netbox.URL+"/graphql/", &httpClient)

	webhookEventBus := make(chan netbox.WebhookEvent, defaultChanSize)

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client := proxmox.NewClient(config.Proxmox.URL,
		proxmox.WithHTTPClient(&insecureHTTPClient),
		proxmox.WithAPIToken(config.Proxmox.TokenID, config.Proxmox.Secret),
	)

	state := state.New(
		state.NewConfig().WithClient(graphqlClient),
		webhookEventBus,
		client,
	)
	api := api.New(
		api.NewConfig().
			WithHost(config.Worker.Host).
			WithPort(config.Worker.Port),
		webhookEventBus,
	)

	// run
	errc := make(chan error, defaultChanSize)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- api.Run()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- state.Run(ctx)
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
	close(webhookEventBus)
	_ = api.Close()
	_ = state.Close()
	wg.Wait()

	return err
}
