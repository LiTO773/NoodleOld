package moodle

import (
	"errors"

	"../db"
)

// SaveUser stores the user in the moodle table
// If it succeeds returns nil, otherwise an error is returned
func SaveUser(url string, username string, password string, token string, location string) (err error) {
	connection, err := db.GetDB()
	if err != nil {
		return
	}

	statement, err := connection.Prepare(`
		INSERT INTO moodles(url, username, password, wstoken, location)
		VALUES (?,?,?,?,?)`)
	if err != nil {
		return
	}

	defer statement.Close()

	res, err := statement.Exec(url, username, password, token, location)
	if err != nil {
		return
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("User wasn't stored")
	}
	return
}

// SaveToken stores the token in the correct place in the DB
// If it succeeds returns nil, otherwise an error is returned
func SaveToken(url string, username string, token string) (err error) {
	connection, err := db.GetDB()
	if err != nil {
		return
	}

	statement, err := connection.Prepare(`
		UPDATE moodles
		SET wstoken=?
		WHERE url=? AND username=?`)
	if err != nil {
		return
	}

	defer statement.Close()

	res, err := statement.Exec(token, url, username)
	if err != nil {
		return
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("Token wasn't stored")
	}
	return
}
