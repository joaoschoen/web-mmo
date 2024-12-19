package pages

import (
	"net/http"
	"web-mmo/modules/templates/t_auth"
	"web-mmo/modules/templates/t_base"
	"web-mmo/modules/templates/t_game"
	"web-mmo/modules/utils/renderer"

	"github.com/labstack/echo/v4"
)

type Page struct {
	Template string
}

func Login(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, t_base.Base(t_auth.Login(), c.Path()))
}

func Hero(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, t_base.GameBase(t_game.Hero(), c.Path()))
}

func World(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, t_base.GameBase(t_game.World(), c.Path()))
}

func Town(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, t_base.GameBase(t_game.Town(), c.Path()))
}

func Dungeon(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, t_base.GameBase(t_game.Dungeon(), c.Path()))
}
