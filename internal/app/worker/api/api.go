package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type api struct {
	config Config
}

func New(config Config) *api {
	return &api{
		config: config,
	}
}

func (api api) Run() error {
	e := echo.New()

	apiRouter := e.Group("/api/v1")
	apiRouter.POST("/update", api.update)

	if err := e.Start(fmt.Sprintf("%s:%d", api.config.host, api.config.port)); err != http.ErrServerClosed {
		return err
	}

	return nil
}
