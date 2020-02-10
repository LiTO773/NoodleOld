package db

import (
	"database/sql"
	"fmt"
	"log"

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

		if err != nil {
			log.Fatalln("Unable to use the DB! " + err)
		}

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
		    courseshash TEXT,
		    unhandlednotifications INTEGER,
		    UNIQUE(url, username)
		);
		CREATE TABLE IF NOT EXISTS courses (
		    moodleid INTEGER UNIQUE,
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    shortname TEXT NOT NULL,
		    fullname TEXT NOT NULL,
		    summary TEXT NOT NULL,
		    downloaded INTEGER NOT NULL,
		    showgrades INTEGER NOT NULL,
		    previoushash TEXT NOT NULL,
		    unhandlednotifications INTEGER,
		    newcourse INTEGER,
		    deletedcourse INTEGER,
		    FOREIGN KEY (moodleid) REFERENCES moodles(id)
		);
		CREATE TABLE IF NOT EXISTS sections (
		    courseid INTEGER UNIQUE,
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    summary TEXT NOT NULL,
		    FOREIGN KEY (courseid) REFERENCES courses(id)
		);
		CREATE TABLE IF NOT EXISTS modules (
		    sectionid INTEGER UNIQUE,
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    modname TEXT NOT NULL,
		    FOREIGN KEY (sectionid) REFERENCES sections(id)
		);
		CREATE TABLE IF NOT EXISTS contents (
		    moduleid INTEGER UNIQUE,
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    indexinmodule INTEGER,
		    type TEXT NOT NULL,
		    filename TEXT NOT NULL,
		    filesize INTEGER NOT NULL,
		    fileurl TEXT NOT NULL,
		    timecreated INTEGER NOT NULL,
		    timemodified INTEGER NOT NULL,
		    author TEXT NOT NULL,
		    license TEXT NOT NULL,
		    FOREIGN KEY (moduleid) REFERENCES modules(id)
		);
		CREATE TABLE IF NOT EXISTS unhandled_moodles (
		    moodleid INTEGER,
		    courseid INTEGER,
		    FOREIGN KEY (moodleid) REFERENCES moodles(id),
		    FOREIGN KEY (courseid) REFERENCES courses(id)
		);
		CREATE TABLE IF NOT EXISTS unhandled_courses (
		    courseid INTEGER,
		    contentid INTEGER,
		    FOREIGN KEY (courseid) REFERENCES courses(id),
		    FOREIGN KEY (contentid) REFERENCES contents(id)
		);`)
	_, err = statement.Exec()
	statement.Close()
	return
}
