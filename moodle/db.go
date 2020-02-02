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

// SearchUser gets one user in the database that matches de parameters given
func SearchUser(url string, username string) (result User, err error) {
	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return
	}

	// Query the DB
	rows, err := connection.Query(`
		SELECT password, wstoken
		FROM moodles
		WHERE url=? AND username=?
	`, url, username)
	if err != nil {
		return
	}
	defer rows.Close()

	// Get and save the first record
	rows.Next()
	err = rows.Scan(&result.Password, &result.Token)

	if err != nil {
		return
	}

	// Add the missing params, this way it will only return a completely filled
	// object when it is successful
	result.Url = url
	result.Username = username

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

// SaveSiteInfo saves the remaining Moodle info
func SaveSiteInfo(info MoodleInfo, url string, username string) (err error) {
	connection, err := db.GetDB()
	if err != nil {
		return
	}

	statement, err := connection.Prepare(`
		UPDATE moodles
		SET sitename=?, firstname=?, lastname=?, lang=?, userid=?, userpictureurl=?
		WHERE url=? AND username=?`)
	if err != nil {
		return
	}

	defer statement.Close()

	res, err := statement.Exec(info.Sitename, info.Firstname, info.Lastname, info.Lang, info.Userid, info.Userpictureurl, url, username)
	if err != nil {
		return
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("Info wasn't stored")
	}
	return
}
