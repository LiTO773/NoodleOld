package moodle

// User stores a user's Moodle url, Username, Password and Token
type User struct {
	URL      string
	Username string
	Password string
	Token    string
}

// InfoMoodle stores the same fields about a Moodle as in the DB
type InfoMoodle struct {
	ID                     int `json:"id"`
	URL                    string
	Username               string
	Password               string
	Token                  string
	Location               string
	Sitename               string `json:"sitename"`
	Firstname              string `json:"firstname"`
	Lastname               string `json:"lastname",omitempty`
	Lang                   string `json:"lang",omitempty`
	Userid                 int    `json:"userid"`
	Userpictureurl         string `json:"userpictureurl",omitempty`
	Previoushash           string
	Unhandlednotifications bool

	ErrorCode string `json:"errorcode",omitempty`
	Message   string `json:"message",omitempty`
}

// Course stores the information currently saved in the DB about a certain course
// request: core_enrol_get_users_courses
type Course struct {
	ID                     int    `json:"id"`
	Shortname              string `json:"shortname"`
	Fullname               string `json:"fullname"`
	Summary                string `json:"summary"`
	Downloaded             bool
	Showgrades             bool `json:"showgrades"`
	Previoushash           string
	Unhandlednotifications bool
	Newcourse              bool
	Deletedcourse          bool
}

// equals Checks if two courses are **completely** equals (not just the id)
func (c Course) equals(c1 Course) bool {
	return c.Shortname == c1.Shortname &&
		c.Fullname == c1.Fullname &&
		c.Summary == c1.Summary &&
		c.Downloaded == c1.Downloaded &&
		c.Showgrades == c1.Showgrades &&
		c.Previoushash == c1.Previoushash &&
		c.Unhandlednotifications == c1.Unhandlednotifications &&
		c.Newcourse == c1.Newcourse &&
		c.Deletedcourse == c1.Deletedcourse
}

// ErrorResponse stores the error code given by the Moodle API
type ErrorResponse struct {
	ErrorCode string `json:"errorcode"`
}
