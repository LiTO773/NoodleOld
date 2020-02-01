package files

import (
	"log"
	"os"
	"runtime"

	errors "../errorcodes"
	"../helpers"
)

// GetSettingsPath returns the correct path to where the application's files are
// stored. If the folder doesn't exist it will be automatically created
// Each OS has a different path:
// Windows: %AppData%\NoodleCore
// macOS: ~/Library/Preferences/NoodleCore
// *nix: ~/.NoodleCore
func GetSettingsPath() (dest string) {
	var currentOS string = runtime.GOOS
	dest = "~/.NoodleCore/"
	if currentOS == "windows" {
		dest = os.Getenv("APPDATA") + "\\NoodleCore\\"
	}

	// Create the folder if it doesn't exist
	err := helpers.CreateLocation(dest)

	if err != nil {
		log.Fatalln("Unable to create the settings folder!")
		log.Fatalln(err)

		os.Exit(int(errors.IOError))
	}

	return dest
}
