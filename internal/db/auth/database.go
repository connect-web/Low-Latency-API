package auth

import (
	"database/sql"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"log"
)

func UsernameExists(username string) (exists bool, error bool) {
	fmt.Println("called usernameExists")
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return false, true
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()
	query := `
	SELECT name 
	FROM users.accs
	where name = $1
	`

	row := client.DB.QueryRow(query, username)

	var databaseUsername string
	err := row.Scan(&databaseUsername)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			fmt.Println("Username does not exist")
			return false, false
		}
		fmt.Println(err.Error())
		fmt.Println("err not nil")
		return false, true
	}

	fmt.Println(databaseUsername)

	exists = databaseUsername != ""
	fmt.Printf("exists: %v\n", exists)
	return exists, err != nil
}

func RegisterUserDatabase(username string, password []byte) (valid bool) {
	if isBytesTooLarge(password) {
		fmt.Printf("Byte protection: Bcrypted Password was over 60 chars!")
		return false
	}

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return false
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()
	query := `
	INSERT INTO users.accs (name, password)
	VALUES ($1, $2)
	`
	_, err := client.DB.Exec(query, username, password)

	return err == nil
}

func LoginGetPassword(username string) (storedPassword []byte, valid bool) {
	var password []byte

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return password, false
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()
	query := `
	SELECT password 
	FROM users.accs
	where name = $1
	`
	row := client.DB.QueryRow(query, username)
	err := row.Scan(&password)

	if len(password) == 0 {
		// 0 byte is invalid result
		return password, false
	}

	return password, err == nil
}
