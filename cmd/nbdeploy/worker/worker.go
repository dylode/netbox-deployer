package worker

import (
	"dylaan.nl/netbox-deployer/internal/app/worker"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker",
		Short: "starts the worker in the foreground",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return worker.Run(args)
		},
	}

	return cmd
}
