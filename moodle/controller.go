package moodle

import (
	errors "../errorcodes"
	files "../files"
	"../helpers"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// NewUser authenticates the user, stores its information in the DB
// and returns the token if it was successfull
// If location is an empty string, it will automatically store to the default
// location
// If it wasn't able to authenticate, then it will return an eror message and
// the corresponding error code
func NewUser(url string, username string, password string, location string) errors.ErrorCode {
	// Check if the URL has the last /
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	// Send and receive the authentication request
	response, err := RequestAuthentication(url, username, password)
	if err != nil {
		return errors.InternetError
	}

	// Check if the request was unsuccessful
	if _, ok := response["error"]; ok {
		return errors.ConvertMoodleError(response["errorcode"].(string))
	}

	var token string = response["token"].(string)

	// The request was successful
	// Get the rest of the information from moodle and store it
	info, err := GetSiteInfo(url, token)
	if err != nil {
		return errors.InternetError
	}

	info.Token = token
	info.URL = url
	info.Username = username
	info.Password = password

	// Check if the location exists, if it doesn't try to create it
	if location == "" {
		// Create the default location
		hash := md5.Sum([]byte(info.Sitename + info.Username))
		folder := hex.EncodeToString(hash[:])
		location = files.GetSettingsPath() + folder
	}
	err = helpers.CreateLocation(location)

	if err != nil {
		return errors.IOError
	}

	info.Location = location

	err = SaveSiteInfo(info, url, username)
	if err != nil {
		return errors.UnableToSaveSiteInfo
	}

	return errors.NoError
}

// LoginUser authenticates a user that already exists in the DB, if the user's
// token is still valid it will login properly, if not it will get a new token
// and login
// It returns the token if it was successful, if not, it returns an error code
func LoginUser(url string, username string) (string, errors.ErrorCode) {
	// Find the user in the DB
	user, err := SearchUser(url, username)

	// Check if the user doesn't exist or the DB crashed
	if err != nil {
		if err.Error() == "sql: Rows are closed" {
			return "", errors.UserDoesNotExist
		}
		return "", errors.DBError
	}

	// The user exists, check if the token is still valid
	info, err := GetSiteInfo(url, user.Token)
	if err != nil {
		return "", errors.InternetError
	}
	// Check for Moodle errors
	if info.ErrorCode != "" {
		if info.ErrorCode != "invalidtoken" {
			return "", errors.ConvertMoodleError(info.ErrorCode)
		}

		// Token isn't valid, get a new one
		response, err := RequestAuthentication(url, username, user.Password)
		if err != nil {
			return "", errors.InternetError
		}

		// Save and return the new token
		var token string = response["token"].(string)

		err = SaveToken(url, username, token)
		if err != nil {
			return "", errors.DBError
		}

		return token, errors.NoError
	}

	// Token is still valid
	return user.Token, errors.NoError
}

// CheckCourses checks if the courses available in Moodle have been added,
// changed or deleted. If that was the case, it will find out which courses were
// altered and handle them appropriately.
// The return values are:
// 1: All the new courses
// 2: All the changed courses
// 3: All the deleted courses
// 4: Error
// NOTE: This function doesn't check if the course's contents have been changed.
func CheckCourses(moodle InfoMoodle) ([]Course, []Course, []Course, errors.ErrorCode) {
	var add, mod, del []Course

	// Get the host courses
	hostCourses, body, errorResponse, err := GetCourses(moodle.URL, moodle.Token, moodle.Userid)

	// Check for errors
	if err != nil {
		return add, mod, del, errors.InternetError
	}
	if errorResponse.ErrorCode != "" {
		if errorResponse.ErrorCode == "invalidtoken" {
			// Token isn't valid anymore
			errorCode := helperUpdateToken(&moodle)
			if errorCode == errors.NoError {
				// A new token was found, try again
				return CheckCourses(moodle)
			}
			return add, mod, del, errorCode
		}
		return add, mod, del, errors.ConvertMoodleError(errorResponse.ErrorCode)
	}

	// Check if there were any changes (https://gist.github.com/sergiotapia/8263278)
	hasher := md5.New()
	hasher.Write(body)
	hash := hex.EncodeToString(hasher.Sum(nil))

	if hash == moodle.Previoushash {
		// Nothing changed
		return add, mod, del, errors.NoError
	}

	// Get the local courses
	localCourses, err := SearchCourses(moodle.ID)
	if err != nil {
		return add, mod, del, errors.DBError
	}

	// Everything is ok, compare
	add, mod, del = helperCourseListComparer(localCourses, hostCourses)

	// Apply changes
	if InsertAddedCourses(moodle.ID, add) != nil || UpdateModifiedCourses(moodle.ID, mod) != nil || UpdateDeletedCourses(moodle.ID, del) != nil {
		return add, mod, del, errors.DBError
	}

	return add, mod, del, errors.NoError
}
