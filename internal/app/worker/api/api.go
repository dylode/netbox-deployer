package api

import (
	"fmt"
	"net/http"

	"dylaan.nl/netbox-deployer/internal/pkg/worker"
	"github.com/labstack/echo"
)

type api struct {
	config Config

	e          *echo.Echo
	updateChan chan worker.Update
}

func New(config Config, updateChan chan worker.Update) *api {
	return &api{
		config: config,

		e:          echo.New(),
		updateChan: updateChan,
	}
}

func (api *api) Run() error {
	apiRouter := api.e.Group("/api/v1")
	apiRouter.POST("/update", api.update)

	if err := api.e.Start(fmt.Sprintf("%s:%d", api.config.host, api.config.port)); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (api *api) Close() error {
	return api.e.Close()
}
