package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	"dylaan.nl/netbox-deployer/internal/app/worker"
	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/Khan/genqlient/graphql"
)

const defaultConfigFilePath = "config.yaml"
const virtualMachineType = "VirtualMachineType"

type node struct {
	name     string
	kind     netbox.TypeKind
	children []node
}

func (n node) String() string {
	var s strings.Builder
	s.WriteString(n.name)
	s.WriteString("\n\r")
	s.WriteString("{ id\n\r")
	for _, field := range n.children {
		s.WriteString(field.String())
		s.WriteString("\n\r")
	}
	s.WriteString("}")
	return s.String()
}

type fieldType = netbox.GetTypeInfoTypeFieldsField

func hasID(fields []fieldType) bool {
	for _, field := range fields {
		if strings.ToLower(field.Name) == "id" {
			return true
		}
	}

	return false
}

func createNode(client graphql.Client, name string, typeName string, typeKind netbox.TypeKind, visited []string) node {
	if slices.Contains(visited, typeName) {
		return node{}
	}
	visited = append(visited, typeName)

	typeInfo, err := netbox.GetTypeInfo(context.Background(), client, typeName)
	if err != nil {
		panic(err)
	}

	if !hasID(typeInfo.Type.Fields) {
		return node{}
	}

	children := []node{}

	if typeKind == netbox.TypeKindObject || typeKind == netbox.TypeKindList {
		for _, field := range typeInfo.Type.Fields {
			if field.Name == "changelog" {
				continue
			}

			switch typeKind {
			case netbox.TypeKindObject:
				children = append(children, createNode(client, field.Name, field.Type.Name, field.Type.Kind, visited))
			case netbox.TypeKindList:
				children = append(children, createNode(client, field.Name, field.Type.OfType.Name, field.Type.Kind, visited))
			}

		}
	}

	return node{
		name:     name,
		kind:     typeKind,
		children: children,
	}
}

//func createNode(client graphql.Client, n nodeInfo, name string, visited []string, kind netbox.TypeKind) node {
//	children := []node{}
//	if slices.Contains(visited, n.Name) {
//		return node{}
//	}
//
//	visited = append(visited, n.Name)
//
//	if kind == netbox.TypeKindObject {
//		for _, field := range n.Fields {
//			if field.Name == "changelog" {
//				continue
//			}
//
//			//typeName := field.Type.Name
//			//if kind == netbox.TypeKindList {
//			//	typeName = field.Type.OfType.Name
//			//}
//
//			//fieldInfo, err := netbox.GetTypeInfo(context.Background(), client, typeName)
//			//if err != nil {
//			//	panic(err)
//			//}
//
//			//if !hasID(fieldInfo.Type.Fields) {
//			//	continue
//			//}
//
//			//children = append(children, createNode(client, fieldInfo.Type, field.Name, visited))
//		}
//	}
//
//	return node{
//		name:     name,
//		kind:     kind,
//		children: children,
//	}
//}

func generate(client graphql.Client) {
	//destination := nodeName("VirtualMachineType")
	//paths := make(map[nodeName]path)
	//visited := []nodeName{}

	root := createNode(client, "virtual_machine_list", virtualMachineType, netbox.TypeKindObject, []string{})
	fmt.Println(root.String())

	//virtualMachineTypeInfo, err := netbox.GetTypeInfo(context.Background(), client, virtualMachineType)
	//if err != nil {
	//	panic(err)
	//}

	//root := createNode(client, virtualMachineTypeInfo.Type, "virtual_machine_list", []string{}, netbox.TypeKindObject)
	//fmt.Println(root.String())

	//var query strings.Builder
	//query.WriteString("query GetVirtualMachines {")
	//query.WriteString("virtual_machine_list {")
	//query.WriteString("id")
	//query.WriteString("}}")

	//fmt.Println(query.String())

	//nodes := schema.GetSchema().Types

	//for _, node := range nodes {
	//	potentialPath := search(client, destination, paths, visited, node)
	//	if potentialPath == nil {
	//		continue
	//	}

	//	paths[node.GetName()] = potentialPath
	//	fmt.Printf("%s => %s", node.GetName(), strings.Join(potentialPath, "->"))
	//}
}

func main() {
	configFilePath := os.Getenv("NBDEPLOY_CONFIG")
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}

	config, err := worker.NewConfigFromPath(configFilePath)
	if err != nil {
		panic(err)
	}

	httpClient := http.Client{}
	graphqlClient := graphql.NewClient(config.Worker.GraphqlURL, &httpClient)

	generate(graphqlClient)
}
