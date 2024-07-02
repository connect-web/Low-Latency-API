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

type MinigameHiscore struct {
	Minigame string
	Count    int
}

// QueryDB executes a query with custom row handling logic.
func (client *DBClient) QueryMinigameHiscores() ([]MinigameHiscore, error) {
	query := `
	SELECT 
	    minigame, 
	    COUNT(DISTINCT id) AS count
	FROM stats.minigame_links
	GROUP BY minigame
	ORDER BY count DESC;
	`
	rows, queryErr := client.DB.Query(query)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return nil, queryErr
	}
	defer rows.Close()

	var results []MinigameHiscore
	for rows.Next() {
		row := MinigameHiscore{}
		if err := rows.Scan(&row.Minigame, &row.Count); err == nil {
			results = append(results, row)
		}
	}
	return results, nil
}

// QueryDB executes a query with custom row handling logic.
func (client *DBClient) QueryMinigameListing(activity string, rowHandler func(*sql.Rows) (model.Player, error)) ([]model.Player, error) {
	query := `
	SELECT
		p.name,
		pl.skills_levels,
		pl.minigames,
		gains.skills_experience,
		gains.minigames
	FROM stats.minigame_links links
	LEFT JOIN PLAYERS P ON P.ID = links.id
	LEFT JOIN player_live PL on PL.playerid = links.id
	LEFT JOIN player_gains gains on gains.playerid = links.id
	WHERE 
	    links.MINIGAME = $1
	ORDER BY links DESC
	LIMIT 133;
	`
	rows, queryErr := client.DB.Query(query, activity)
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
