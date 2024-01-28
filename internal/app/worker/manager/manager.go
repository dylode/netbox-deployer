package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/luthermonson/go-proxmox"
	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashset"
)

const defaultSetSize = 1_000

type netboxClient interface {
	GetManagingVirtualMachines(context.Context) (map[netbox.ModelID]netbox.VirtualMachine, error)
}

func ModelIDHash(modelID netbox.ModelID) uint64 {
	return g.HashInt(int(modelID))
}

type manager struct {
	sync.Mutex

	webhookEventBus       <-chan netbox.WebhookEvent
	updatables            *hashset.Set[netbox.ModelID]
	netboxVirtualMachines map[netbox.ModelID]netbox.VirtualMachine
	pveClient             *proxmox.Client
	netboxClient          netboxClient
}

func New(webhookEventBus <-chan netbox.WebhookEvent, pveClient *proxmox.Client, netboxClient netboxClient) *manager {
	return &manager{
		webhookEventBus: webhookEventBus,
		updatables:      hashset.New[netbox.ModelID](defaultSetSize, g.Equals, ModelIDHash),
		//	netboxVirtualMachines: make(map[netbox.ModelID]netbox.VirtualMachine),
		pveClient:    pveClient,
		netboxClient: netboxClient,
	}
}

func (r *manager) sync(ctx context.Context) error {
	defer r.Unlock()
	r.Lock()

	if r.updatables.Size() == 0 && len(r.netboxVirtualMachines) != 0 {
		return nil
	}

	if err := r.updateNetboxVirtualMachines(ctx); err != nil {
		return err
	}

	fmt.Println(r.netboxVirtualMachines)

	// TODO: loop through updatables, check validaty, update proxmox accordingly
	//r.updatables.Each(func(id netbox.ModelID) {
	//	fmt.Printf("updating virtual machine %d\n\r", id)

	//	vm := r.netboxVirtualMachines[id]
	//	vmFlat := vm.Flat()

	//	if vm.PveID != nil {

	//		node, err := r.pveClient.Node(ctx, "pve01")
	//		if err != nil {
	//			panic(err)
	//		}

	//		pvevm, err := node.VirtualMachine(ctx, *vm.PveID)
	//		if err != nil {
	//			panic(err)
	//		}

	//		existingTags := strings.Split(pvevm.Tags, ";")
	//		for _, t := range existingTags {
	//			if !slices.Contains(vmFlat.Tags, t) {
	//				pvevm.RemoveTag(ctx, t)
	//			}
	//		}

	//		for _, t := range vmFlat.Tags {
	//			if !slices.Contains(pvevm.VirtualMachineConfig.TagsSlice, t) {
	//				pvevm.AddTag(ctx, t)
	//			}
	//		}

	//	}
	//})

	r.updatables.Clear()

	return nil
}

func (r *manager) updateNetboxVirtualMachines(ctx context.Context) error {
	netboxVirtualMachines, err := r.netboxClient.GetManagingVirtualMachines(ctx)
	if err != nil {
		return err
	}

	r.netboxVirtualMachines = netboxVirtualMachines

	return nil
}

func (r *manager) Run(ctx context.Context) error {
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

func (r *manager) process(event netbox.WebhookEvent) {
	//defer r.Unlock()

	fmt.Printf("%s %s %d\n\r", event.EventType, event.ModelName, event.ModelID)

	//if event.EventType == netbox.EventCreated && slices.Contains(netbox.AllModelNames(), event.ModelName) {
	//	r.Lock()
	//	for id := range r.netboxVirtualMachines {
	//		r.updatables.Put(id)
	//	}
	//	return
	//}

	//wg := sync.WaitGroup{}
	//wg.Add(len(r.netboxVirtualMachines))
	//updatables := make(chan netbox.ModelID, len(r.netboxVirtualMachines))

	//for id, vm := range r.netboxVirtualMachines {
	//	// TODO: there is something better for this
	//	myVM := vm
	//	myID := id
	//	go func() {
	//		defer wg.Done()
	//		if myVM.HasRelation(event.ModelName, event.ModelID) {
	//			updatables <- myID
	//		}
	//	}()
	//}

	//wg.Wait()
	//close(updatables)

	//r.Lock()
	//for id := range updatables {
	//	r.updatables.Put(id)
	//}
}

func (r *manager) Close() error {
	return nil
}