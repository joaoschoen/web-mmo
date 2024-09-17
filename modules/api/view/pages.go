package view

import (
	"web-mmo/modules/api/controller/pages"

	"github.com/labstack/echo/v4"
)

func PagesRouter(server *echo.Echo) {
	group := server.Group("/")
	group.GET("", pages.World)
	group.GET("world", pages.World)
	group.GET("hero", pages.Hero)
	group.GET("town", pages.Town)
	group.GET("dungeon", pages.Dungeon)
}
