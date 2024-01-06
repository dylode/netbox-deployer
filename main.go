package main

import (
	"os"

	"dylaan.nl/netbox-deployer/cmd/nbdeploy"
)

//go:generate go run github.com/Khan/genqlient
//go:generate go run ./internal/app/worker/resolver/scheme/generate.go

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
