package helpers

// ErrorCode number that corresponds to a known error
type ErrorCode int

// Only the controllers use these error codes
const (
	NoError          ErrorCode = 0 // gg
	InternetError    ErrorCode = 1 // Error sending/receiving a request
	DBError          ErrorCode = 2 // Error reading/writing to the DB
	UserDoesNotExist ErrorCode = 6 // User isn't registered in the DB
	IOError          ErrorCode = 3 // Error reading/writing files
	MoodleError      ErrorCode = 4 // Unknown error given by Moodle
	WebServicesError ErrorCode = 5 // Error given when the Web services are deactivated
)
