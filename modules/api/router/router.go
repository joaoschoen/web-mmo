package router

import (
	// "web-mmo/view"

	"web-mmo/modules/api/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	view.PagesRouter(e)
	view.UserRouter(e)
}
