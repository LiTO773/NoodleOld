package helpers

// ErrorCode number that corresponds to a known error
// If the number is negative, it is a warning
type ErrorCode int

const (
	NoError              ErrorCode = 0 // gg
	InternetError        ErrorCode = 1 // Error sending/receiving a request
	DBError              ErrorCode = 2 // Error reading/writing to the DB
	IOError              ErrorCode = 3 // Error reading/writing files
	MoodleError          ErrorCode = 4 // Unknown error given by Moodle
	WebServicesError     ErrorCode = 5 // Error given when the Web services are deactivated
	UserDoesNotExist     ErrorCode = 6 // User isn't registered in the DB
	UnableToSaveSiteInfo ErrorCode = 7 // Site info could not be saved in the DB
	InvalidTokenError    ErrorCode = 8 // The current token isn't valid anymore
	InvalidLoginError    ErrorCode = 9 // Username/Password are incorrect
)

// ConvertMoodleError Receives a Moodle Error Code and converts to the
// appropriate ErrorCode
func ConvertMoodleError(err string) ErrorCode {
	switch err {
	case "invalidtoken":
		return InvalidTokenError
	case "invalidlogin":
		return InvalidLoginError
	case "enablewsdescription":
		return WebServicesError
	}
	return MoodleError
}
