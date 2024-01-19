package state

import "dylaan.nl/netbox-deployer/internal/pkg/netbox"

func walk(current relation, targetModelName netbox.ModelName, targetModelID netbox.ModelID) bool {
	if current.getModelName() == targetModelName && current.getModelID() == targetModelID {
		return true
	}

	for _, relation := range current.getRelations() {
		if walk(relation, targetModelName, targetModelID) {
			return true
		}
	}

	return false
}

func (vm virtualMachine) hasRelation(targetModelName netbox.ModelName, targetModelID netbox.ModelID) bool {
	return walk(&vm, targetModelName, targetModelID)
}
