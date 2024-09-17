package pages

import (
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	// BODY
	return c.Render(200, "index", nil)
}
