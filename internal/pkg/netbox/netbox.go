package netbox

import (
	"context"
	"fmt"

	gonetbox "github.com/netbox-community/go-netbox/v3"
)

type netbox struct {
	*gonetbox.APIClient
}

func New(url string, token string) *netbox {
	c := gonetbox.NewAPIClientFor(url, token)
	return &netbox{
		APIClient: c,
	}
}

func (netbox netbox) getManagingCluster(ctx context.Context) ([]gonetbox.Cluster, error) {
	// TODO: get rid of limit
	clusterList, _, err := netbox.
		VirtualizationAPI.
		VirtualizationClustersList(ctx).
		Tag([]string{"managed-by-netbox-deployer"}).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}
	return clusterList.Results, nil
}

func (netbox netbox) getManagingVirtualMachines(ctx context.Context, clusters []gonetbox.Cluster) ([]gonetbox.VirtualMachineWithConfigContext, error) {
	clusterIDs := make([]*int32, 0, len(clusters))
	for _, cluster := range clusters {
		id := cluster.GetId()
		clusterIDs = append(clusterIDs, &id)
	}

	// TODO: get rid of limit
	virtualMachines, _, err := netbox.
		VirtualizationAPI.
		VirtualizationVirtualMachinesList(ctx).
		ClusterId(clusterIDs).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}

	return virtualMachines.GetResults(), nil
}

func (netbox netbox) getInterfacesForVM(ctx context.Context, vmID int32) ([]gonetbox.VMInterface, error) {
	interfaces, _, err := netbox.
		VirtualizationAPI.
		VirtualizationInterfacesList(ctx).
		VirtualMachineId([]int32{vmID}).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}

	return interfaces.GetResults(), nil
}

func (netbox netbox) getDisksForVM(ctx context.Context, vmID int32) ([]gonetbox.VirtualDisk, error) {
	disks, _, err := netbox.
		VirtualizationAPI.
		VirtualizationVirtualDisksList(ctx).
		VirtualMachineId([]int32{vmID}).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}

	return disks.GetResults(), nil
}

func (netbox netbox) getIPAddressesForInterface(ctx context.Context, interfaceID int32) ([]gonetbox.IPAddress, error) {
	ipAddresses, _, err := netbox.
		IpamAPI.
		IpamIpAddressesList(ctx).
		VminterfaceId([]int32{interfaceID}).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}

	return ipAddresses.GetResults(), nil
}

func (netbox netbox) SetVirtualMachinePlanned(ctx context.Context, vmID int32) error {
	status := gonetbox.MODULESTATUSVALUE_PLANNED
	_, _, err := netbox.
		VirtualizationAPI.
		VirtualizationVirtualMachinesPartialUpdate(ctx, vmID).
		PatchedWritableVirtualMachineWithConfigContextRequest(gonetbox.PatchedWritableVirtualMachineWithConfigContextRequest{
			Status: &status,
		}).
		Execute()
	return err
}

func (netbox netbox) WriteVirtualMachineJournal(ctx context.Context, vmID int32, comment string) error {
	_, _, err := netbox.
		ExtrasAPI.
		ExtrasJournalEntriesCreate(ctx).
		WritableJournalEntryRequest(*gonetbox.NewWritableJournalEntryRequest("virtualization.virtualmachine", int64(vmID), comment)).
		Execute()
	return err
}

func (netbox netbox) GetManagingVirtualMachines(ctx context.Context) (map[ModelID]VirtualMachine, error) {
	clusters, err := netbox.getManagingCluster(ctx)
	if err != nil {
		return nil, err
	}

	virtualMachines, err := netbox.getManagingVirtualMachines(ctx, clusters)
	if err != nil {
		fmt.Println("hier")
		return nil, err
	}

	vms := make(map[ModelID]VirtualMachine)
	for _, vm := range virtualMachines {
		vms[ModelID(vm.GetId())], err = netbox.newVirtualMachine(ctx, vm)
		if err != nil {
			return nil, err
		}
	}

	return vms, nil
}
