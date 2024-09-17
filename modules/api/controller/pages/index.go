package pages

import (
	"net/http"
	"web-mmo/modules/templates"
	"web-mmo/modules/utils/renderer"

	"github.com/labstack/echo/v4"
)

type Page struct {
	Template string
}

func Hero(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, templates.Base(templates.Hero(), c.Path()))
}

func World(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, templates.Base(templates.World(), c.Path()))
}

func Town(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, templates.Base(templates.Town(), c.Path()))
}

func Dungeon(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, templates.Base(templates.Dungeon(), c.Path()))
}
