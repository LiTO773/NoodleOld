package db

import (
	"database/sql"

	"../files"
)

const fileName = "stuff.db"

// currentConn stores the current connection with the SQLite database
var currentConn *sql.DB

// GetDB returns or creates a connection with the application's SQLite database.
// If the database doesn't exist, a new one will be created
func GetDB() (err error) {
	if currentConn == nil {
		// Start the connection
		currentConn, err = sql.Open("sqlite3", files.GetSettingsPath()+fileName)

		// Test query to check if the DB exists
		_, err = currentConn.Query("SELECT id FROM moodles")
		if err != nil {
			// Populate the DB
			err = populateDB()
		}
	}
	return
}

// populateDB creates the required tables for the application to work
func populateDB() (err error) {
	statement, err := currentConn.Prepare(`CREATE TABLE IF NOT EXISTS moodles (
		id INTEGER PRIMARY KEY,
		url TEXT,
		username TEXT,
		wstoken TEXT,
		userid INTEGER,
		location TEXT,
		UNIQUE(url, username)
		)`)
	_, err = statement.Exec()
	return
}
