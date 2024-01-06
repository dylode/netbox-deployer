package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dylaan.nl/netbox-deployer/internal/pkg/worker"
	"github.com/labstack/echo"
)

func (api api) update(c echo.Context) error {
	req := c.Request()
	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var update worker.Update
	if err = json.Unmarshal(body, &update); err != nil {
		return err
	}

	fmt.Println(update)

	return c.NoContent(http.StatusOK)
}
