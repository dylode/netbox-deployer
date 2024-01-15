package state

import (
	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

type vmData = netbox.GetVirtualMachinesVirtual_machine_listVirtualMachineType
type tag = netbox.GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType
type vmInterface = netbox.GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType
type ipAddress = netbox.GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType
type taggedVlan = netbox.GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType

func newVirtualMachine(vmData vmData) virtualMachine {
	vm := &virtualMachine{
		baseRelation: newBaseRelation(vmData.Id),
	}

	initTags(vm, vmData.Tags)
	initVirtualMachineInterfaces(vm, vmData.Interfaces)

	return *vm
}

func initTags[T relation](parent T, tags []tag) {
	for _, tag := range tags {
		parent.addRelation(&tagRelation{
			baseRelation: newBaseRelation(tag.Id),
			name:         tag.Name,
		})
	}
}

func initVirtualMachineInterfaces[T relation](parent T, vmInterfaces []vmInterface) {
	for _, vmInterface := range vmInterfaces {
		interf := &virtualMachineInterfaceRelation{
			baseRelation: newBaseRelation(vmInterface.Id),
			name:         vmInterface.Name,
		}

		initIPAddresses(interf, vmInterface.Ip_addresses)
		initTaggedVlans(interf, vmInterface.Tagged_vlans)

		parent.addRelation(interf)
	}
}

func initIPAddresses[T relation](parent T, ipAddresses []ipAddress) {
	for _, ipAddress := range ipAddresses {
		parent.addRelation(&ipAddressRelation{
			baseRelation: newBaseRelation(ipAddress.Id),
			address:      ipAddress.Address,
		})
	}
}

func initTaggedVlans[T relation](parent T, vlans []taggedVlan) {
	for _, vlan := range vlans {
		parent.addRelation(&taggedVlanRelation{
			baseRelation: newBaseRelation(vlan.Id),
			vid:          vlan.Vid,
		})
	}
}
