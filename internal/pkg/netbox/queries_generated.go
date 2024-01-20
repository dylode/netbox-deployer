// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package netbox

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetVirtualMachinesResponse is returned by GetVirtualMachines on success.
type GetVirtualMachinesResponse struct {
	Virtual_machine_list []GetVirtualMachinesVirtual_machine_listVirtualMachineType `json:"virtual_machine_list"`
}

// GetVirtual_machine_list returns GetVirtualMachinesResponse.Virtual_machine_list, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesResponse) GetVirtual_machine_list() []GetVirtualMachinesVirtual_machine_listVirtualMachineType {
	return v.Virtual_machine_list
}

// GetVirtualMachinesVirtual_machine_listVirtualMachineType includes the requested fields of the GraphQL type VirtualMachineType.
type GetVirtualMachinesVirtual_machine_listVirtualMachineType struct {
	Id            string                                                                              `json:"id"`
	Tags          []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType               `json:"tags"`
	Interfaces    []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType `json:"interfaces"`
	Custom_fields map[string]any                                                                      `json:"custom_fields"`
}

// GetId returns GetVirtualMachinesVirtual_machine_listVirtualMachineType.Id, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineType) GetId() string { return v.Id }

// GetTags returns GetVirtualMachinesVirtual_machine_listVirtualMachineType.Tags, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineType) GetTags() []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType {
	return v.Tags
}

// GetInterfaces returns GetVirtualMachinesVirtual_machine_listVirtualMachineType.Interfaces, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineType) GetInterfaces() []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType {
	return v.Interfaces
}

// GetCustom_fields returns GetVirtualMachinesVirtual_machine_listVirtualMachineType.Custom_fields, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineType) GetCustom_fields() map[string]any {
	return v.Custom_fields
}

// GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType includes the requested fields of the GraphQL type VMInterfaceType.
type GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType struct {
	Id           string                                                                                                       `json:"id"`
	Name         string                                                                                                       `json:"name"`
	Ip_addresses []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType `json:"ip_addresses"`
	Tagged_vlans []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType      `json:"tagged_vlans"`
}

// GetId returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType.Id, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType) GetId() string {
	return v.Id
}

// GetName returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType.Name, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType) GetName() string {
	return v.Name
}

// GetIp_addresses returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType.Ip_addresses, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType) GetIp_addresses() []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType {
	return v.Ip_addresses
}

// GetTagged_vlans returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType.Tagged_vlans, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceType) GetTagged_vlans() []GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType {
	return v.Tagged_vlans
}

// GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType includes the requested fields of the GraphQL type IPAddressType.
type GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType struct {
	Id string `json:"id"`
	// IPv4 or IPv6 address (with mask)
	Address string `json:"address"`
}

// GetId returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType.Id, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType) GetId() string {
	return v.Id
}

// GetAddress returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType.Address, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeIp_addressesIPAddressType) GetAddress() string {
	return v.Address
}

// GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType includes the requested fields of the GraphQL type VLANType.
type GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType struct {
	Id string `json:"id"`
	// Numeric VLAN ID (1-4094)
	Vid int `json:"vid"`
}

// GetId returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType.Id, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType) GetId() string {
	return v.Id
}

// GetVid returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType.Vid, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeInterfacesVMInterfaceTypeTagged_vlansVLANType) GetVid() int {
	return v.Vid
}

// GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType includes the requested fields of the GraphQL type TagType.
type GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetId returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType.Id, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType) GetId() string {
	return v.Id
}

// GetName returns GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType.Name, and is useful for accessing the field via an interface.
func (v *GetVirtualMachinesVirtual_machine_listVirtualMachineTypeTagsTagType) GetName() string {
	return v.Name
}

// The query or mutation executed by GetVirtualMachines.
const GetVirtualMachines_Operation = `
query GetVirtualMachines {
	virtual_machine_list {
		id
		tags {
			id
			name
		}
		interfaces {
			id
			name
			ip_addresses {
				id
				address
			}
			tagged_vlans {
				id
				vid
			}
		}
		custom_fields
	}
}
`

func GetVirtualMachines(
	ctx context.Context,
	client graphql.Client,
) (*GetVirtualMachinesResponse, error) {
	req := &graphql.Request{
		OpName: "GetVirtualMachines",
		Query:  GetVirtualMachines_Operation,
	}
	var err error

	var data GetVirtualMachinesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
