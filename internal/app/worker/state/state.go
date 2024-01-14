package state

import (
	"context"
	"fmt"
	"strconv"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
)

const workerUpdateChanBuffer = 1_000

type relation interface {
	getModel() netbox.Model
	getModelID() netbox.ModelID
}

type virtualMachine struct {
	id        string
	relations []relation
}

type tagRelation struct {
	id   netbox.ModelID
	name string
}

func (tagRelation) getModel() netbox.Model {
	return netbox.Model("tag")
}

func (t tagRelation) getModelID() netbox.ModelID {
	return t.id
}

type state struct {
	config Config

	updateChan      chan netbox.Update
	virtualMachines []virtualMachine
}

func New(config Config) *state {
	return &state{
		config: config,

		updateChan:      make(chan netbox.Update, workerUpdateChanBuffer),
		virtualMachines: []virtualMachine{},
	}
}

func (r state) GetUpdateChan() chan netbox.Update {
	return r.updateChan
}

func (r *state) Run(ctx context.Context) error {
	allVirtualMachinesRequest, err := netbox.GetVirtualMachines(ctx, r.config.client)
	if err != nil {
		return err
	}

	for _, vm := range allVirtualMachinesRequest.Virtual_machine_list {
		relations := []relation{}

		for _, tag := range vm.Tags {
			id, err := strconv.Atoi(tag.Id)
			if err != nil {
				return err
			}

			relations = append(relations, tagRelation{
				id:   netbox.ModelID(id),
				name: tag.Name,
			})
		}

		r.virtualMachines = append(r.virtualMachines, virtualMachine{
			id:        vm.Id,
			relations: relations,
		})
	}

	for update := range r.updateChan {
		for _, vm := range r.virtualMachines {
			for _, relation := range vm.relations {
				if relation.getModel() == update.Model && relation.getModelID() == update.ID {
					fmt.Println(vm)
				}
			}
		}
	}

	return nil
}

func (r state) Close() error {
	close(r.updateChan)
	return nil
}
