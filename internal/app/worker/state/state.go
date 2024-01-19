package state

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashset"
)

const defaultSetSize = 1_000

func ModelIDHash(modelID netbox.ModelID) uint64 {
	return g.HashInt(int(modelID))
}

type state struct {
	sync.Mutex
	config Config

	webhookEventBus       <-chan netbox.WebhookEvent
	updatables            *hashset.Set[netbox.ModelID]
	netboxVirtualMachines map[netbox.ModelID]netbox.VirtualMachine
}

func New(config Config, webhookEventBus <-chan netbox.WebhookEvent) *state {
	return &state{
		config: config,

		webhookEventBus:       webhookEventBus,
		updatables:            hashset.New[netbox.ModelID](defaultSetSize, g.Equals, ModelIDHash),
		netboxVirtualMachines: make(map[netbox.ModelID]netbox.VirtualMachine),
	}
}

func (r *state) sync(ctx context.Context) error {
	defer r.Unlock()
	r.Lock()

	if r.updatables.Size() == 0 && len(r.netboxVirtualMachines) != 0 {
		return nil
	}

	if err := r.updateNetboxVirtualMachines(ctx); err != nil {
		return err
	}

	// TODO: loop through updatables, check validaty, update proxmox accordingly
	r.updatables.Each(func(id netbox.ModelID) {
		fmt.Printf("updating virtual machine %d\n\r", id)
	})

	r.updatables.Clear()

	return nil
}

func (r *state) updateNetboxVirtualMachines(ctx context.Context) error {
	clear(r.netboxVirtualMachines)

	netboxVirtualMachines, err := netbox.GetVirtualMachines(ctx, r.config.client)
	if err != nil {
		return err
	}

	for _, vm := range netboxVirtualMachines.Virtual_machine_list {
		netboxVM := netbox.NewVirtualMachine(vm)
		r.netboxVirtualMachines[netboxVM.ID] = netboxVM
	}

	return nil
}

func (r *state) Run(ctx context.Context) error {
	err := r.sync(ctx)
	if err != nil {
		return err
	}

	// TODO: make configurable
	syncTicker := time.NewTicker(time.Second)

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case <-syncTicker.C:
			r.sync(ctx)
		case event := <-r.webhookEventBus:
			r.process(event)
		}
	}

	return nil
}

func (r *state) process(event netbox.WebhookEvent) {
	defer r.Unlock()

	fmt.Printf("%s %s %d\n\r", event.EventType, event.ModelName, event.ModelID)

	if event.EventType == netbox.EventCreated && slices.Contains(netbox.AllModelNames(), event.ModelName) {
		r.Lock()
		for id := range r.netboxVirtualMachines {
			r.updatables.Put(id)
		}
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(r.netboxVirtualMachines))
	updatables := make(chan netbox.ModelID, len(r.netboxVirtualMachines))

	for id, vm := range r.netboxVirtualMachines {
		// TODO: there is something better for this
		myVM := vm
		myID := id
		go func() {
			defer wg.Done()
			if myVM.HasRelation(event.ModelName, event.ModelID) {
				updatables <- myID
			}
		}()
	}

	wg.Wait()
	close(updatables)

	r.Lock()
	for id := range updatables {
		r.updatables.Put(id)
	}
}

func (r *state) Close() error {
	return nil
}
