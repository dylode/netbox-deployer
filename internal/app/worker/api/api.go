package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type api struct {
	config Config

	e *echo.Echo
}

func New(config Config) *api {
	return &api{
		config: config,
	}
}

func (api *api) Run() error {
	api.e = echo.New()

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
