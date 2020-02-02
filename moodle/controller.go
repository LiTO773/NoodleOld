package moodle

import (
	errors "../errorcodes"
	"../helpers"
)

// NewUser authenticates the user, stores its information in the DB
// and returns the token if it was successfull
// If it wasn't able to authenticate, then it will return an eror message and
// the corresponding error code
func NewUser(url string, username string, password string, location string) (string, errors.ErrorCode) {
	// Check if the location exists, if it doesn't try to create it
	err := helpers.CreateLocation(location)

	if err != nil {
		return "", errors.IOError
	}

	// Send and receive the authentication request
	response, err := RequestAuthentication(url, username, password)
	if err != nil {
		return "", errors.InternetError
	}

	// Check if the request was unsuccessful
	if _, ok := response["error"]; ok {
		// Check if Web Services are disabled in Moodle
		if response["errorcode"].(string) == "enablewsdescription" {
			return "", errors.WebServicesError
		}
		// Something else happened
		return response["error"].(string), errors.MoodleError
	}

	var token string = response["token"].(string)

	// The request was successful, so store the token in the DB
	err = SaveUser(url, username, password, token, location)
	if err != nil {
		return "", errors.DBError
	}

	// Get the rest of the information from moodle and store it
	info, err := GetSiteInfo(url, token)
	if err != nil {
		return "", errors.InternetError
	}

	err = SaveSiteInfo(info, url, username)
	if err != nil {
		return "", errors.UnableToSaveSiteInfo
	}

	return token, errors.NoError
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
			return "", errors.MoodleError
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
