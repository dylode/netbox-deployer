package netbox

import (
	"context"

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

func (netbox netbox) getManagingClusterIDS(ctx context.Context) ([]gonetbox.Cluster, error) {
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
	clusterIDs := make([]int32, 0, len(clusters))
	for _, cluster := range clusters {
		clusterIDs = append(clusterIDs, cluster.GetId())
	}

	// TODO: get rid of limit
	virtualMachines, _, err := netbox.
		VirtualizationAPI.
		VirtualizationVirtualMachinesList(ctx).
		ClusterGroupId(clusterIDs).
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
		InterfaceId([]int32{interfaceID}).
		Limit(100).
		Execute()
	if err != nil {
		return nil, err
	}

	return ipAddresses.GetResults(), nil
}

func (netbox netbox) GetManagingVirtualMachines(ctx context.Context) (map[ModelID]VirtualMachine, error) {
	clusters, err := netbox.getManagingClusterIDS(ctx)
	if err != nil {
		return nil, err
	}

	virtualMachines, err := netbox.getManagingVirtualMachines(ctx, clusters)
	if err != nil {
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
