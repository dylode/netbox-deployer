package state

import (
	"context"
	"fmt"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

type state struct {
	config Config

	webhookEventBus <-chan netbox.WebhookEvent

	netboxVirtualMachines map[netbox.ModelID]netbox.VirtualMachine
}

func New(config Config, webhookEventBus <-chan netbox.WebhookEvent) *state {
	return &state{
		config: config,

		webhookEventBus:       webhookEventBus,
		netboxVirtualMachines: make(map[netbox.ModelID]netbox.VirtualMachine),
	}
}

func (r *state) initState(ctx context.Context) error {
	allVirtualMachinesRequest, err := netbox.GetVirtualMachines(ctx, r.config.client)
	if err != nil {
		return err
	}

	for _, vm := range allVirtualMachinesRequest.Virtual_machine_list {
		netboxVM := netbox.NewVirtualMachine(vm)
		r.netboxVirtualMachines[netboxVM.ID] = netboxVM
	}

	return nil
}

func (r *state) Run(ctx context.Context) error {
	err := r.initState(ctx)
	if err != nil {
		return err
	}

	for event := range r.webhookEventBus {
		r.process(event)
	}

	return nil
}

func (r state) process(event netbox.WebhookEvent) {
	for id, vm := range r.netboxVirtualMachines {
		if vm.HasRelation(event.ModelName, event.ModelID) {
			fmt.Println(id)
		}
	}
}

func (r state) Close() error {
	return nil
}
