package main

import (
	"fmt"
	"log"
	"os"
	"web-mmo/modules/api/router"
	"web-mmo/modules/utils/database"
	"web-mmo/modules/utils/environment"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Instantiate API
	e := echo.New()

	// Environment config
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
		os.Exit(1)
	}
	if err := environment.Validate(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	e.Static("/css", "static/css")

	router.InitRoutes(e)

	// Database config
	// Init job to connect to db
	go database.CheckDBConnection()

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
