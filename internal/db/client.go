package db

import (
	"database/sql"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/model"
	_ "github.com/lib/pq"
	"log"
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

// Connect establishes a connection to the database.
func (client *DBClient) Connect() error {
	user := os.Getenv("lowLatencyUser")
	password := os.Getenv("lowLatencyPassword")

	host := os.Getenv("lowLatencyHost")
	port := os.Getenv("lowLatencyPort")

	dbname := os.Getenv("lowLatencyDatabase")

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

// QueryDB executes a query with custom row handling logic.
func (client *DBClient) QueryDBSimplePlayers(query string, params []interface{}, rowHandler func(*sql.Rows) (model.SimplePlayer, error)) ([]model.SimplePlayer, error) {
	fmt.Println("Sending Query")
	rows, queryErr := client.DB.Query(query, params...)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return nil, queryErr
	}
	defer rows.Close()

	var results []model.SimplePlayer
	for rows.Next() {
		result, rowParseError := rowHandler(rows)
		if rowParseError != nil {
			log.Println("Error handling row:", rowParseError.Error())
			continue // or return, depending on how critical an error in one row is
		}
		results = append(results, result)
	}
	return results, nil
}

// QueryDB executes a query with custom row handling logic.
func (client *DBClient) QueryDBPlayers(query string, params []interface{}, rowHandler func(*sql.Rows) (model.Player, error)) ([]model.Player, error) {
	fmt.Println("Sending Query")
	rows, queryErr := client.DB.Query(query, params...)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return nil, queryErr
	}
	defer rows.Close()

	var results []model.Player
	for rows.Next() {
		result, rowParseError := rowHandler(rows)
		if rowParseError != nil {
			log.Println("Error handling row:", rowParseError.Error())
			continue // or return, depending on how critical an error in one row is
		}
		results = append(results, result)
	}
	return results, nil
}
