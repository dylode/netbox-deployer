package netbox

func walk(current relation, targetModelName ModelName, targetModelID ModelID) bool {
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

func (vm VirtualMachine) HasRelation(targetModelName ModelName, targetModelID ModelID) bool {
	return walk(&vm, targetModelName, targetModelID)
}
