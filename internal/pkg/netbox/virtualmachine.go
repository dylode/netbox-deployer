package netbox

import (
	"context"

	gonetbox "github.com/netbox-community/go-netbox/v3"
)

func newVMComponent[T any](id int32, component T) VirtualMachineComponent[T] {
	return VirtualMachineComponent[T]{
		ID:   ModelID(id),
		Data: component,
	}
}

func (netbox netbox) newVirtualMachine(ctx context.Context, root gonetbox.VirtualMachineWithConfigContext) (VirtualMachine, error) {
	vm := VirtualMachine{
		Status: string(root.Status.GetValue()),
		CPUs:   uint(root.GetVcpus()),
		Memory: uint64(root.GetMemory()),
	}

	for _, tag := range root.Tags {
		vm.Tags = append(vm.Tags, newVMComponent[string](tag.GetId(), tag.GetName()))
	}

	vmInterfaces, err := netbox.getInterfacesForVM(ctx, root.GetId())
	if err != nil {
		return vm, err
	}

	for _, vmInterface := range vmInterfaces {
		ipAddresses, err := netbox.getIPAddressesForInterface(ctx, vmInterface.GetId())
		if err != nil {
			return vm, err
		}

		ipAddressesInterface := []VirtualMachineComponent[InterfaceIPAddress]{}
		for _, ipAddress := range ipAddresses {
			ipAddressesInterface = append(ipAddressesInterface, newVMComponent[InterfaceIPAddress](ipAddress.GetId(), InterfaceIPAddress{
				Address: ipAddress.GetAddress(),
			}))
		}

		vlan := vmInterface.GetUntaggedVlan()

		vm.Interfaces = append(vm.Interfaces, newVMComponent[VirtualMachineInterface](vmInterface.GetId(), VirtualMachineInterface{
			VLAN: newVMComponent[InterfaceVLAN](vlan.GetId(), InterfaceVLAN{
				VID: vlan.GetVid(),
			}),
			MacAddress:  vmInterface.GetMacAddress(),
			IPAddresses: ipAddressesInterface,
		}))
	}

	vmDisks, err := netbox.getDisksForVM(ctx, root.GetId())
	if err != nil {
		return vm, err
	}

	for _, disk := range vmDisks {
		vm.Disks = append(vm.Disks, newVMComponent[VirtualMachineDisk](disk.GetId(), VirtualMachineDisk{
			Name: disk.GetName(),
			Size: uint64(disk.GetSize()),
		}))
	}

	return vm, nil
}
