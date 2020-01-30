package moodle

import "fmt"

// AuthenticateUser authenticates the user, stores it's information in the DB
// and returns the token if it was successfull
// If it wasn't able to authenticate it returns an eror message and an error
// code
func AuthenticateUser(hostname string, username string, password string) (bool, error) {
	// Send and receive the authentication request
	response, err := RequestAuthentication(hostname, username, password)
	if err != nil {
		return false, err
	}

	// Check if the request was successful
	fmt.Println(response)

	return true, nil
}
