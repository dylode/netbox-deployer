package api

import (
	"encoding/json"
	"io"
	"net/http"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/labstack/echo"
)

func (api api) update(c echo.Context) error {
	req := c.Request()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var event netbox.WebhookEvent
	if err = json.Unmarshal(body, &event); err != nil {
		return err
	}

	api.webhookEventBus <- event

	return c.NoContent(http.StatusOK)
}
