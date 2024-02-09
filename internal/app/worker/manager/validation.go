package manager

import (
	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

func (r *manager) validationCheck(vm netbox.VirtualMachine) []string {
	errors := []string{}

	if vm.CPUs == 0 {
		errors = append(errors, "CPU count cannot be zero")
	}

	if vm.Memory == 0 {
		errors = append(errors, "Memory cannot be zero")
	}

	return errors
}
