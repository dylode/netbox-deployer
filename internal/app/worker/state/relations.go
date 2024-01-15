package state

import (
	"strconv"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

type relation interface {
	getModel() netbox.Model
	getModelID() netbox.ModelID
	addRelation(relation)
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

func (virtualMachine) getModel() netbox.Model {
	return netbox.Model("virtualmachine")
}

type tagRelation struct {
	baseRelation
	name string
}

func (tagRelation) getModel() netbox.Model {
	return netbox.Model("tag")
}

type virtualMachineInterfaceRelation struct {
	baseRelation
	name string
}

func (virtualMachineInterfaceRelation) getModel() netbox.Model {
	return netbox.Model("interface")
}

type ipAddressRelation struct {
	baseRelation
	address string
}

func (ipAddressRelation) getModel() netbox.Model {
	return netbox.Model("ipaddress")
}

type taggedVlanRelation struct {
	baseRelation
	vid int
}

func (taggedVlanRelation) getModel() netbox.Model {
	return netbox.Model("taggedvlan")
}
