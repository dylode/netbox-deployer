package worker

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dylaan.nl/netbox-deployer/internal/app/worker/api"
)

func Run(config Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	apiConfig := api.NewConfig().
		WithHost(config.Worker.Host).
		WithPort(config.Worker.Port)

	api := api.New(apiConfig)

	errc := make(chan error, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- api.Run()
	}()

	var err error
	select {
	case <-ctx.Done():
		fmt.Println("closing due to interrupt")
	case err = <-errc:
		fmt.Println("closing due to error")
	}

	cancel()
	_ = api.Close()
	wg.Wait()

	return err
}
