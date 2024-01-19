package state

import (
	"context"
	"sync"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashset"
	"golang.org/x/exp/maps"
)

type state struct {
	sync.Mutex
	config Config

	webhookEventBus       <-chan netbox.WebhookEvent
	updatables            *hashset.Set[int]
	netboxVirtualMachines map[netbox.ModelID]netbox.VirtualMachine
}

func New(config Config, webhookEventBus <-chan netbox.WebhookEvent) *state {
	return &state{
		config: config,

		webhookEventBus:       webhookEventBus,
		updatables:            hashset.New[int](10, g.Equals, g.HashInt),
		netboxVirtualMachines: make(map[netbox.ModelID]netbox.VirtualMachine),
	}
}

func (r *state) sync(ctx context.Context) error {
	defer r.Unlock()
	r.Lock()

	maps.Clear(r.netboxVirtualMachines)

	allVirtualMachinesRequest, err := netbox.GetVirtualMachines(ctx, r.config.client)
	if err != nil {
		return err
	}

	for _, vm := range allVirtualMachinesRequest.Virtual_machine_list {
		netboxVM := netbox.NewVirtualMachine(vm)
		r.netboxVirtualMachines[netboxVM.ID] = netboxVM
	}

	// TODO: loop through updatables, check validaty, update proxmox accordingly

	return nil
}

func (r *state) Run(ctx context.Context) error {
	err := r.sync(ctx)
	if err != nil {
		return err
	}

	for event := range r.webhookEventBus {
		r.process(event)
	}

	return nil
}

func (r *state) process(event netbox.WebhookEvent) {
	for id, vm := range r.netboxVirtualMachines {
		// TODO, in case of CREATE only check model name
		if vm.HasRelation(event.ModelName, event.ModelID) {
			r.updatables.Put(int(id))
		}
	}
}

func (r *state) Close() error {
	return nil
}
