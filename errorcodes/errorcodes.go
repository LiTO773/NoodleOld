package helpers

type ErrorCode int

// Only the controllers use these error codes
const (
	NoError       ErrorCode = 0 // gg
	InternetError ErrorCode = 1 // Error sending/receiving a request
	DBError       ErrorCode = 2 // Error reading/writing to the DB
	IOError       ErrorCode = 3 // error reading/writing files
	MoodleError   ErrorCode = 4 // Error given by Moodle
)
