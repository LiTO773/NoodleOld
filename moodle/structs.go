package moodle

// User stores a user's Moodle url, Username, Password and Token
type User struct {
	Url      string
	Username string
	Password string
	Token    string
}

// MoodleInfo stores part of the information of the request: webservice/rest/server.php?moodlewsrestformat=json
type MoodleInfo struct {
	Sitename       string `json:"sitename",omitempty`
	Firstname      string `json:"firstname",omitempty`
	Lastname       string `json:"lastname",omitempty`
	Lang           string `json:"lang",omitempty`
	Userid         int    `json:"userid",omitempty`
	Userpictureurl string `json:"userpictureurl",omitempty`

	ErrorCode string `json:"errorcode",omitempty`
	Message   string `json:"message",omitempty`
}
