package files

import (
	"log"
	"os"
	"runtime"

	"../helpers"
)

// GetSettingsPath returns the correct path to where the application's files are
// stored. If the folder doesn't exist it will be automatically created
// Each OS has a different path:
// Windows: %AppData%\Noodle
// macOS: ~/Library/Preferences/Noodle
// *nix: ~/.Noodle
func GetSettingsPath() (dest string) {
	var currentOS string = runtime.GOOS
	dest = "~/.Noodle/"
	if currentOS == "windows" {
		dest = os.Getenv("APPDATA") + "\\Noodle\\"
	}

	// Create the folder if it doesn't exist
	err := helpers.CreateLocation(dest)

	if err != nil {
		log.Fatalln("Unable to create the settings folder! " + err.Error())
	}

	return dest
}
