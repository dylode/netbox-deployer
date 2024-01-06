package resolver

import (
	"fmt"

	"dylaan.nl/netbox-deployer/internal/pkg/worker"
)

const workerUpdateChanBuffer = 1_000

type resolver struct {
	config Config

	updateChan chan worker.Update
}

func New(config Config) *resolver {
	return &resolver{
		config: config,

		updateChan: make(chan worker.Update, workerUpdateChanBuffer),
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
