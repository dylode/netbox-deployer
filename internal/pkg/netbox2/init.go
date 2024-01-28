package netbox

type vmData = GetVirtualMachinesVirtual_machine_listVirtualMachineType
type tag = GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType
type vmInterface = GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType
type ipAddress = GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType
type taggedVlan = GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType

func NewVirtualMachine(vmData vmData) VirtualMachine {
	pveID, ok := vmData.Custom_fields["nb_pve_id"].(float64)
	var pveID2 *int
	if ok {
		pveIDInt := int(pveID)
		pveID2 = &pveIDInt
	}
	vm := &VirtualMachine{
		baseRelation: newBaseRelation(vmData.Id),
		PveID:        pveID2,
	}

	initTags(vm, vmData.Tags)
	initVirtualMachineInterfaces(vm, vmData.Interfaces)

	return *vm
}

func initTags(parent relation, tags []tag) {
	for _, tag := range tags {
		parent.addRelation(&TagRelation{
			baseRelation: newBaseRelation(tag.Id),
			Name:         tag.Name,
		})
	}
}

func initVirtualMachineInterfaces(parent relation, vmInterfaces []vmInterface) {
	for _, vmInterface := range vmInterfaces {
		interf := &VirtualMachineInterfaceRelation{
			baseRelation: newBaseRelation(vmInterface.Id),
			Name:         vmInterface.Name,
		}

		initIPAddresses(interf, vmInterface.Ip_addresses)
		initTaggedVlans(interf, vmInterface.Tagged_vlans)

		parent.addRelation(interf)
	}
}

func initIPAddresses(parent relation, ipAddresses []ipAddress) {
	for _, ipAddress := range ipAddresses {
		parent.addRelation(&IPAddressRelation{
			baseRelation: newBaseRelation(ipAddress.Id),
			Address:      ipAddress.Address,
		})
	}
}

func initTaggedVlans(parent relation, vlans []taggedVlan) {
	for _, vlan := range vlans {
		parent.addRelation(&TaggedVlanRelation{
			baseRelation: newBaseRelation(vlan.Id),
			VID:          vlan.Vid,
		})
	}
}
