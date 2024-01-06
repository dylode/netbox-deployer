// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package netbox

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// An enumeration.
type IpamIPAddressStatusChoices string

const (
	// Active
	IpamIPAddressStatusChoicesActive IpamIPAddressStatusChoices = "ACTIVE"
	// Reserved
	IpamIPAddressStatusChoicesReserved IpamIPAddressStatusChoices = "RESERVED"
	// Deprecated
	IpamIPAddressStatusChoicesDeprecated IpamIPAddressStatusChoices = "DEPRECATED"
	// DHCP
	IpamIPAddressStatusChoicesDhcp IpamIPAddressStatusChoices = "DHCP"
	// SLAAC
	IpamIPAddressStatusChoicesSlaac IpamIPAddressStatusChoices = "SLAAC"
)

// testIp_addressIPAddressType includes the requested fields of the GraphQL type IPAddressType.
type testIp_addressIPAddressType struct {
	Id string `json:"id"`
	// IPv4 or IPv6 address (with mask)
	Address     string                                      `json:"address"`
	Description string                                      `json:"description"`
	Vrf         testIp_addressIPAddressTypeVrfVRFType       `json:"vrf"`
	Tenant      testIp_addressIPAddressTypeTenantTenantType `json:"tenant"`
	// The operational status of this IP
	Status IpamIPAddressStatusChoices `json:"status"`
}

// GetId returns testIp_addressIPAddressType.Id, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetId() string { return v.Id }

// GetAddress returns testIp_addressIPAddressType.Address, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetAddress() string { return v.Address }

// GetDescription returns testIp_addressIPAddressType.Description, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetDescription() string { return v.Description }

// GetVrf returns testIp_addressIPAddressType.Vrf, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetVrf() testIp_addressIPAddressTypeVrfVRFType { return v.Vrf }

// GetTenant returns testIp_addressIPAddressType.Tenant, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetTenant() testIp_addressIPAddressTypeTenantTenantType {
	return v.Tenant
}

// GetStatus returns testIp_addressIPAddressType.Status, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressType) GetStatus() IpamIPAddressStatusChoices { return v.Status }

// testIp_addressIPAddressTypeTenantTenantType includes the requested fields of the GraphQL type TenantType.
type testIp_addressIPAddressTypeTenantTenantType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetId returns testIp_addressIPAddressTypeTenantTenantType.Id, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressTypeTenantTenantType) GetId() string { return v.Id }

// GetName returns testIp_addressIPAddressTypeTenantTenantType.Name, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressTypeTenantTenantType) GetName() string { return v.Name }

// testIp_addressIPAddressTypeVrfVRFType includes the requested fields of the GraphQL type VRFType.
type testIp_addressIPAddressTypeVrfVRFType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetId returns testIp_addressIPAddressTypeVrfVRFType.Id, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressTypeVrfVRFType) GetId() string { return v.Id }

// GetName returns testIp_addressIPAddressTypeVrfVRFType.Name, and is useful for accessing the field via an interface.
func (v *testIp_addressIPAddressTypeVrfVRFType) GetName() string { return v.Name }

// testResponse is returned by test on success.
type testResponse struct {
	Ip_address testIp_addressIPAddressType `json:"ip_address"`
}

// GetIp_address returns testResponse.Ip_address, and is useful for accessing the field via an interface.
func (v *testResponse) GetIp_address() testIp_addressIPAddressType { return v.Ip_address }

// The query or mutation executed by test.
const test_Operation = `
query test {
	ip_address(id: 2) {
		id
		address
		description
		vrf {
			id
			name
		}
		tenant {
			id
			name
		}
		status
	}
}
`

func test(
	ctx context.Context,
	client graphql.Client,
) (*testResponse, error) {
	req := &graphql.Request{
		OpName: "test",
		Query:  test_Operation,
	}
	var err error

	var data testResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}