package moodle

import (
	errors "../errorcodes"
	"../helpers"
)

// NewUser authenticates the user, stores its information in the DB
// and returns the token if it was successfull
// If it wasn't able to authenticate, then it will return an eror message and
// the corresponding error code
func NewUser(hostname string, username string, password string, location string) (string, errors.ErrorCode) {
	// Check if the location exists, if it doesn't try to create it
	err := helpers.CreateLocation(location)

	if err != nil {
		return "", errors.IOError
	}

	// Send and receive the authentication request
	response, err := RequestAuthentication(hostname, username, password)
	if err != nil {
		return "", errors.InternetError
	}

	// Check if the request was unsuccessful
	if _, ok := response["error"]; ok {
		return response["error"].(string), errors.MoodleError
	}

	var token string = response["token"].(string)

	// The request was successful, so store the token in the DB
	err = SaveUser(hostname, username, password, token, location)

	if err != nil {
		return "", errors.DBError
	}

	return token, errors.NoError
}
