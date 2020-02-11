package moodle

import (
	"../db"
	"errors"
)

// SaveUser stores the user in the moodle table
// If it succeeds returns nil, otherwise an error is returned
func SaveUser(url string, username string, password string, token string, location string) error {
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	statement, err := connection.Prepare(`
		INSERT INTO moodles(url, username, password, wstoken, location)
		VALUES (?,?,?,?,?)`)
	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(url, username, password, token, location)
	if err != nil {
		return err
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("User wasn't stored")
	}
	return err
}

// SearchUser gets one user in the database that matches de parameters given
func SearchUser(url string, username string) (User, error) {
	var result User

	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return result, err
	}

	// Query the DB
	rows, err := connection.Query(`
		SELECT password, wstoken
		FROM moodles
		WHERE url=? AND username=?
	`, url, username)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	// Get and save the first record
	rows.Next()
	err = rows.Scan(&result.Password, &result.Token)

	if err != nil {
		return result, err
	}

	// Add the missing params, this way it will only return a completely filled
	// object when it is successful
	result.URL = url
	result.Username = username

	return result, err
}

// SaveToken stores the token in the correct place in the DB
// If it succeeds returns nil, otherwise an error is returned
func SaveToken(url string, username string, token string) error {
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	statement, err := connection.Prepare(`
		UPDATE moodles
		SET wstoken=?
		WHERE url=? AND username=?`)
	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(token, url, username)
	if err != nil {
		return err
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("Token wasn't stored")
	}
	return err
}

// SaveSiteInfo saves the remaining Moodle info
func SaveSiteInfo(info InfoMoodle, url string, username string) error {
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	statement, err := connection.Prepare(`
		UPDATE moodles
		SET sitename=?, firstname=?, lastname=?, lang=?, userid=?, userpictureurl=?
		WHERE url=? AND username=?`)
	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(info.Sitename, info.Firstname, info.Lastname, info.Lang, info.Userid, info.Userpictureurl, url, username)
	if err != nil {
		return err
	} else if num, err := res.RowsAffected(); num == 0 || err != nil {
		err = errors.New("Info wasn't stored")
	}
	return err
}

// SearchMoodle returns the Moodle that corresponds to the arguments passed
func SearchMoodle(url string, username string) (InfoMoodle, error) {
	var result InfoMoodle
	result.URL = url
	result.Username = username

	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return result, err
	}

	// Query the DB
	rows, err := connection.Query(`
		SELECT id, password, wstoken, location, sitename, firstname, lastname, lang, userid, userpictureurl, previoushash, unhandlednotifications
		FROM moodles
		WHERE url=? AND username=?
	`, url, username)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	// Get and save to the struct
	rows.Next()
	err = rows.Scan(
		&result.ID,
		&result.Password,
		&result.Token,
		&result.Location,
		&result.Sitename,
		&result.Firstname,
		&result.Lastname,
		&result.Lang,
		&result.Userid,
		&result.Userpictureurl,
		&result.Previoushash,
		&result.Unhandlednotifications)

	return result, err
}

// SearchCourses returns all the courses that exist locally
func SearchCourses(moodleid int) ([]Course, error) {
	var result []Course
	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return result, err
	}

	// Query the DB
	rows, err := connection.Query(`
		SELECT id, shortname, fullname, summary, downloaded, showgrades, previoushash, unhandlednotifications, newcourse, deletedcourse
		FROM courses
		WHERE moodleid=?
	`, moodleid)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	// Get and save to the struct
	for rows.Next() {
		var c Course
		err = rows.Scan(
			&c.ID,
			&c.Shortname,
			&c.Fullname,
			&c.Summary,
			&c.Downloaded,
			&c.Showgrades,
			&c.Previoushash,
			&c.Unhandlednotifications,
			&c.Newcourse,
			&c.Deletedcourse)
		if err != nil {
			return result, err
		}
		result = append(result, c)
	}

	return result, err
}

// InsertAddedCourses inserts into the DB all the new courses.
// NOTE: It doesn't change the unhandlednotifications status
func InsertAddedCourses(moodleid int, courses []Course) error {
	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	// https://golang.org/pkg/database/sql/#Tx.Query
	tx, err := connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare(`
		INSERT INTO courses(moodleid, id, shortname, fullname, summary, downloaded, showgrades, previoushash, unhandlednotifications, newcourse, deletedcourse)
		VALUES (?,?,?,?,?,?,?,?,?,1,0)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range courses {
		_, err := stmt.Exec(
			moodleid,
			c.ID,
			c.Shortname,
			c.Fullname,
			c.Summary,
			c.Downloaded,
			c.Showgrades,
			c.Previoushash,
			c.Unhandlednotifications)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateModifiedCourses changes the properties of already existing courses.
// NOTE: It doesn't change the unhandlednotifications status
func UpdateModifiedCourses(moodleid int, courses []Course) error {
	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	tx, err := connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare(`
		UPDATE courses
		SET shortname=?, fullname=?, summary=?, showgrades=?
		WHERE moodleid=? AND id=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range courses {
		_, err := stmt.Exec(
			c.Shortname,
			c.Fullname,
			c.Summary,
			c.Showgrades,
			moodleid,
			c.ID)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateModifiedCourses changes the properties of already existing courses.
// NOTE: It doesn't change the unhandlednotifications status
func UpdateDeletedCourses(moodleid int, courses []Course) error {
	// Get the connection
	connection, err := db.GetDB()
	if err != nil {
		return err
	}

	tx, err := connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare(`
		UPDATE courses
		SET deletedcourse=1
		WHERE moodleid=? AND id=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range courses {
		_, err := stmt.Exec(moodleid, c.ID)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
