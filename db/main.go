package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
		os.Exit(1)
	}
	db_conn, err := conn()
	if err != nil {
		log.Fatalln(err.Error())
	}
	
	println("hello world")
}

func conn() (*sql.DB, error) {
	conn_string := os.Getenv("DATABASE_URL")
	if conn_string == "" {
		log.Fatalf("Couldn't load DATABASE_URL environment variable properly")
	}

	var err error
	dbConn, err := sql.Open("postgres", conn_string)

	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		return nil, err
	}
	return dbConn, nil
}
