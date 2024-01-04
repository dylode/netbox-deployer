package nbdeploy

import (
	"dylaan.nl/netbox-deployer/cmd/nbdeploy/worker"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   "nbdeploy",
		Short: "nbdeploy uses Netbox as source-of-truth to create virtual machines on Proxmox",
	}

	cmd.AddCommand(worker.NewCommand())

	return cmd
}
