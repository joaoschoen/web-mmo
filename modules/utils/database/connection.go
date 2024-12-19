package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func InitDB() error {
	conn_string := os.Getenv("DATABASE_URL")
	if conn_string == "" {
		log.Fatalf("Couldn't load DATABASE_URL environment variable properly")
	}

	var err error
	dbConn, err = sql.Open("postgres", conn_string)

	if err != nil {
		return err
	}
	dbConn.SetMaxOpenConns(20)

	if err := dbConn.Ping(); err != nil {
		return err
	}

	log.Println("Database connection established")
	return nil
}

func CheckDBConnection() {
	for {
		ctx := context.Background()
		var err error
		if dbConn == nil {
			err := recreateDBPool()
			if err != nil {
				log.Printf("Unable to acquire connection: %v", err)
				time.Sleep(30 * time.Second) // Check every 30 seconds
				continue
			}
		}
		conn, err := dbConn.Conn(ctx)
		if err != nil {
			log.Printf("Unable to acquire connection: %v", err)
			time.Sleep(30 * time.Second) // Check every 30 seconds
			continue
		}

		// Ping the database
		err = conn.PingContext(ctx)
		if err != nil {
			log.Printf("Database ping failed: %v", err)
			// Recreate the connection pool
			err := recreateDBPool()
			if err != nil {
				log.Printf("Unable to acquire connection: %v", err)
				time.Sleep(30 * time.Second) // Check every 30 seconds
				continue
			}
		} else {
			log.Println("Database connection is healthy")
		}
		time.Sleep(30 * time.Second) // Check every 30 seconds
	}
}

func GetConnection(ctx context.Context) (*sql.Conn, error) {
	return dbConn.Conn(ctx)
}

func recreateDBPool() error {
	// Close the existing pool if it exists
	if dbConn != nil {
		dbConn.Close()
	}

	return InitDB() // Reinitialize the connection pool
}
