package files

import (
	"fmt"
	"os"
	"runtime"
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
	} else if currentOS == "darwin" {
		dest = "~/Library/Preferences/NoodleCore/"
	}

	// Check if it exists
	// https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		// Doesn't exist, so create the folder
		err = os.Mkdir(dest, os.ModePerm)
		fmt.Println(err)
	}

	return dest
}
