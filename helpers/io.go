package helpers

import (
	"os"
)

// CreateLocation checks if the given path exists, if not, it tries to create it
func CreateLocation(location string) error {
	_, err := os.Stat(location)

	// Check if it exists
	// https://gist.github.com/mattes/d13e273314c3b3ade33f
	if os.IsNotExist(err) {
		// Doesn't exist, so create the folder
		err = os.Mkdir(location, os.ModePerm)
	}

	return err
}
