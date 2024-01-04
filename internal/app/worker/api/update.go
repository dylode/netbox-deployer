package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func (api api) update(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return c.NoContent(http.StatusOK)
}
