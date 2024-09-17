package view

import (
	"web-mmo/modules/api/controller/pages"

	"github.com/labstack/echo/v4"
)

func PagesRouter(server *echo.Echo) {
	group := server.Group("/")
	group.GET("", pages.Index)
}
