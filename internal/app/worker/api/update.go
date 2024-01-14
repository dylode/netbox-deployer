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

	var update netbox.Update
	if err = json.Unmarshal(body, &update); err != nil {
		return err
	}

	api.updateChan <- update

	return c.NoContent(http.StatusOK)
}
