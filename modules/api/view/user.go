package view

import (
	// "API-ECHO/controller/admin"

	"web-mmo/modules/api/controller/user"

	"github.com/labstack/echo/v4"
)

func UserRouter(server *echo.Echo) {
	group := server.Group("/auth")
	group.POST("/login", user.Login)
	group.POST("/register", user.Register)
}
