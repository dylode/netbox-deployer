package netbox

import (
	"strconv"
)

type relation interface {
	getModelName() ModelName
	getModelID() ModelID
	addRelation(relation)
	getRelations() []relation
}

type baseRelation struct {
	ID        ModelID
	relations []relation
}

func (b baseRelation) getModelID() ModelID {
	return b.ID
}

func (b *baseRelation) addRelation(r relation) {
	b.relations = append(b.relations, r)
}

func (b baseRelation) getRelations() []relation {
	return b.relations
}

func newBaseRelation(id string) baseRelation {
	modelID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	return baseRelation{
		ID:        ModelID(modelID),
		relations: []relation{},
	}
}

type VirtualMachine struct {
	baseRelation
}

func (VirtualMachine) getModelName() ModelName {
	return ModelName("virtualmachine")
}

type TagRelation struct {
	baseRelation
	Name string
}

func (TagRelation) getModelName() ModelName {
	return ModelName("tag")
}

type VirtualMachineInterfaceRelation struct {
	baseRelation
	Name string
}

func (VirtualMachineInterfaceRelation) getModelName() ModelName {
	return ModelName("vminterface")
}

type IPAddressRelation struct {
	baseRelation
	Address string
}

func (IPAddressRelation) getModelName() ModelName {
	return ModelName("ipaddress")
}

type TaggedVlanRelation struct {
	baseRelation
	VID int
}

func (TaggedVlanRelation) getModelName() ModelName {
	return ModelName("vlan")
}
