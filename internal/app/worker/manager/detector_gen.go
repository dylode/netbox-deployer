// Code generated by "gen.go"; DO NOT EDIT.

package manager

import "dylaan.nl/netbox-deployer/internal/pkg/netbox"

var allModelNames []netbox.ModelName

func init() {
	allModelNames = []netbox.ModelName {
		netbox.ModelName("virtualmachine"),
		netbox.ModelName("tag"),
		netbox.ModelName("vminterface"),
		netbox.ModelName("vmdisk"),
	}
}