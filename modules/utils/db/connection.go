package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

// TODO REMOVE THIS
// const conn_string string = "host=localhost port=5432 database='web-mmo' user='admin' password='admin' pool_min_conns=5 pool_max_conns=20 pool_max_conn_lifetime=15s ssl=false"
const conn_string string = "postgres://admin:admin@localhost:5432/web-mmo?sslmode=disable"

func InitDB() error {
	var err error
	dbConn, err = sql.Open("postgres", conn_string)
	dbConn.SetMaxOpenConns(20)

	if err != nil {
		log.Printf("Failed to connect to the database: %v", err.Error())
		return err
	}

	if dbConn.Ping() != nil {
		log.Printf("Failed to connect to the database: %v", err.Error())
	}

	log.Println("Database connection established")
	return nil
}

func CheckDBConnection() {
	for {
		ctx := context.Background()
		conn, err := dbConn.Conn(ctx)
		if err != nil {
			log.Printf("Unable to acquire connection: %v", err)
			continue
		}

		// Ping the database
		err = conn.PingContext(ctx)
		if err != nil {
			log.Printf("Database ping failed: %v", err)
			// Recreate the connection pool
			recreateDBPool()
		} else {
			log.Println("Database connection is healthy")
		}
		time.Sleep(30 * time.Second) // Check every 30 seconds
	}
}

func GetConnection(ctx context.Context) (*sql.Conn, error) {
	return dbConn.Conn(ctx)
}

func recreateDBPool() {
	// Close the existing pool if it exists
	if dbConn != nil {
		dbConn.Close()
	}

	InitDB() // Reinitialize the connection pool
}
