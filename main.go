package main

import (
	"fmt"
	"log"
	"web-mmo/modules/api/router"
	"web-mmo/modules/utils/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init API
	e := echo.New()

	e.Static("/css", "static/css")

	router.InitRoutes(e)

	// Database config
	// Initialize the DB connection pool
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}

	go db.CheckDBConnection()

	// Middleware stack
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// Initialize server
	e.Logger.Fatal(e.Start(fmt.Sprint(":", 3000)))
}
