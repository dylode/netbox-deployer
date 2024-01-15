package state

import (
	"context"
	"fmt"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

const workerUpdateChanBuffer = 1_000

type state struct {
	config Config

	updateChan      chan netbox.Update
	virtualMachines []virtualMachine
}

func New(config Config) *state {
	return &state{
		config: config,

		updateChan:      make(chan netbox.Update, workerUpdateChanBuffer),
		virtualMachines: []virtualMachine{},
	}
}

func (r state) GetUpdateChan() chan netbox.Update {
	return r.updateChan
}

func (r *state) initState(ctx context.Context) error {
	allVirtualMachinesRequest, err := netbox.GetVirtualMachines(ctx, r.config.client)
	if err != nil {
		return err
	}

	for _, vm := range allVirtualMachinesRequest.Virtual_machine_list {
		r.virtualMachines = append(r.virtualMachines, newVirtualMachine(vm))
	}

	return nil
}

func (r *state) Run(ctx context.Context) error {
	err := r.initState(ctx)
	if err != nil {
		return err
	}
	fmt.Println(r.virtualMachines)

	for update := range r.updateChan {
		for _, vm := range r.virtualMachines {
			for _, relation := range vm.relations {
				if relation.getModel() == update.Model && relation.getModelID() == update.ID {
					fmt.Println(vm)
				}
			}
		}
	}

	return nil
}

func (r state) Close() error {
	close(r.updateChan)
	return nil
}
