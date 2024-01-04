package worker

import (
	"dylaan.nl/netbox-deployer/internal/app/worker"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var configFilePath string

	cmd := &cobra.Command{
		Use:   "worker",
		Short: "starts the worker in the foreground",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := worker.NewConfigFromPath(configFilePath)
			if err != nil {
				return err
			}

			return worker.Run(config)
		},
	}

	cmd.Flags().StringVar(&configFilePath, "config", "config.yaml", "path to config file")

	return cmd
}
