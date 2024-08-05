package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// DBClient represents a client to the database with a connection status.
type DBClient struct {
	DB              *sql.DB
	Connected       bool
	ConnectionError error
}

// NewDBClient initializes a new database client.
func NewDBClient() *DBClient {
	return &DBClient{
		Connected: false,
	}
}

func GetUrl() string {
	envVar := os.Getenv("siteonline")

	var user, password, host, port, dbname string

	if envVar == "site" {
		user = os.Getenv("lowLatencyUser")
		password = os.Getenv("lowLatencyPassword")
		host = os.Getenv("lowLatencyHost")
		port = os.Getenv("lowLatencyPort")
		dbname = os.Getenv("lowLatencyDatabase")
	} else {
		// LOCALHOST connect to server credentials
		user = os.Getenv("lowLatencyWebUser")
		password = os.Getenv("lowLatencyWebPassword")
		host = os.Getenv("lowLatencyWebHost")
		port = os.Getenv("lowLatencyWebPort")
		dbname = os.Getenv("lowLatencyWebDatabase")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}

// Connect establishes a connection to the database.
func (client *DBClient) Connect() error {

	envVar := os.Getenv("siteonline")

	var user, password, host, port, dbname string

	if envVar == "site" {
		user = os.Getenv("lowLatencyUser")
		password = os.Getenv("lowLatencyPassword")
		host = os.Getenv("lowLatencyHost")
		port = os.Getenv("lowLatencyPort")
		dbname = os.Getenv("lowLatencyDatabase")
	} else {
		// LOCALHOST connect to server credentials
		user = os.Getenv("lowLatencyWebUser")
		password = os.Getenv("lowLatencyWebPassword")
		host = os.Getenv("lowLatencyWebHost")
		port = os.Getenv("lowLatencyWebPort")
		dbname = os.Getenv("lowLatencyWebDatabase")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		client.ConnectionError = err
		return err
	}

	// Try to make a connection
	err = db.Ping()
	if err != nil {
		client.ConnectionError = err
		return err
	}

	client.DB = db
	client.Connected = true
	client.ConnectionError = nil
	return nil
}

// Close terminates the connection to the database.
func (client *DBClient) Close() error {
	if client.DB != nil {
		return client.DB.Close()
	}
	return nil
}
