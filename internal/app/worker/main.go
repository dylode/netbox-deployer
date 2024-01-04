package worker

import "dylaan.nl/netbox-deployer/internal/app/worker/api"

func Run(args []string) error {
	apiConfig := api.NewConfig().
		WithHost("10.10.10.1").
		WithPort(8080)

	api := api.New(apiConfig)

	return api.Run()
}
