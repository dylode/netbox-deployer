package main

import (
	"os"

	"dylaan.nl/netbox-deployer/cmd/nbdeploy"

	_ "github.com/Khan/genqlient"
)

//go:generate go run github.com/Khan/genqlient

func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(args []string) error {
	cmd := nbdeploy.NewCommand()
	cmd.SetArgs(args)
	return cmd.Execute()
}
