package db

import (
	"database/sql"
	"fmt"

	"../files"
)

// FileName stores the SQLite file name
const FileName = "stuff.db"

// currentConn stores the current connection with the SQLite database
var currentConn *sql.DB

// GetDB returns or creates a connection with the application's SQLite database.
// If the database doesn't exist, a new one will be created
func GetDB() (*sql.DB, error) {
	if currentConn == nil {
		fmt.Println("Connection opened!")
		// Start the connection
		var err error
		currentConn, err = sql.Open("sqlite3", files.GetSettingsPath()+FileName)

		// Test query to check if the DB has the expected tables
		_, err = currentConn.Exec("SELECT id FROM moodles")
		if err != nil {
			// Populate the DB
			err = populateDB()
		}
		return currentConn, err
	}
	return currentConn, nil
}

// populateDB creates the required tables for the application to work
func populateDB() (err error) {
	statement, err := currentConn.Prepare(`
		CREATE TABLE IF NOT EXISTS moodles (
			id INTEGER PRIMARY KEY
			           AUTOINCREMENT,
			url TEXT NOT NULL
			         CHECK ((url GLOB 'http://*/') OR (url GLOB 'https://*/')),
			username TEXT NOT NULL
			              CHECK (LENGTH(username) > 0),
			password TEXT NOT NULL
			              CHECK (LENGTH(password) > 0),
			wstoken TEXT,
			location TEXT NOT NULL
			              CHECK (LENGTH(location) > 0),
			sitename TEXT,
			firstname TEXT,
			lastname TEXT,
			lang TEXT,
			userid TEXT,
			userpictureurl TEXT,
			UNIQUE(url, username)
		);`)
	_, err = statement.Exec()
	statement.Close()
	return
}
