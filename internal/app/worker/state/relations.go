package state

import (
	"strconv"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

type relation interface {
	getModelName() netbox.ModelName
	getModelID() netbox.ModelID
	addRelation(relation)
	getRelations() []relation
}

type baseRelation struct {
	id        netbox.ModelID
	relations []relation
}

func (b baseRelation) getModelID() netbox.ModelID {
	return b.id
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
		id:        netbox.ModelID(modelID),
		relations: []relation{},
	}
}

type virtualMachine struct {
	baseRelation
}

func (virtualMachine) getModelName() netbox.ModelName {
	return netbox.ModelName("virtualmachine")
}

type tagRelation struct {
	baseRelation
	name string
}

func (tagRelation) getModelName() netbox.ModelName {
	return netbox.ModelName("tag")
}

type virtualMachineInterfaceRelation struct {
	baseRelation
	name string
}

func (virtualMachineInterfaceRelation) getModelName() netbox.ModelName {
	return netbox.ModelName("vminterface")
}

type ipAddressRelation struct {
	baseRelation
	address string
}

func (ipAddressRelation) getModelName() netbox.ModelName {
	return netbox.ModelName("ipaddress")
}

type taggedVlanRelation struct {
	baseRelation
	vid int
}

func (taggedVlanRelation) getModelName() netbox.ModelName {
	return netbox.ModelName("vlan")
}
