package moodle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// RequestAuthentication requests Moodle for a user's token
// Resquest: <host>/login/token.php?service=moodle_mobile_app
// Body: username: username, password: password
// Returns the user's token in the Moodle or an error code if Moodle Web
// Services aren't allowed in the server
// NOTE: If a Moodle service is https:// and the user types http://, Moodle will
// return an error saying that the username wasn't provided!
func RequestAuthentication(host string, username string, password string) (map[string]interface{}, error) {
	var result map[string]interface{}

	resp, err := http.PostForm(host+"login/token.php?service=moodle_mobile_app", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		log.Println(err)
		return result, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return result, err
	}

	// Get the content
	err = json.Unmarshal([]byte(body), &result)

	return result, err
}

// GetSiteInfo requests the information about the user and the moodle
// Request: <host>/webservice/rest/server.php?moodlewsrestformat=json
// Body: wstoken: token, wsfunction: core_webservice_get_site_info
// Returns a InfoMoodle object if the request was successful (even if an error
// was received) or an error if it wasn't able to connect to the server
func GetSiteInfo(host string, token string) (InfoMoodle, error) {
	var result InfoMoodle

	resp, err := http.PostForm(host+"webservice/rest/server.php?moodlewsrestformat=json", url.Values{
		"wstoken":    {token},
		"wsfunction": {"core_webservice_get_site_info"},
	})
	if err != nil {
		log.Println(err)
		return result, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	// Get the content
	err = json.Unmarshal([]byte(body), &result)

	return result, err
}

// GetCourses requests the information of all the courses available to the user
// Request: <host>/webservice/rest/server.php?moodlewsrestformat=json
// Body: wstoken: token, wsfunction: core_enrol_get_users_courses, userid: userid
// Returns a list of Courses and an error (if one was given during the operation)
// result: Returns the courses found
// response: Returns the response received (later to be hashed)
// er: Moodle returned an error
// err: The operation was unsuccessful
func GetCourses(host string, token string, userid int) (result []Course, body []byte, er ErrorResponse, err error) {
	var resp *http.Response
	resp, err = http.PostForm(host+"webservice/rest/server.php?moodlewsrestformat=json", url.Values{
		"wstoken":    {token},
		"wsfunction": {"core_enrol_get_users_courses"},
		"userid":     {strconv.Itoa(userid)},
	})
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(body))

	// Get the content
	err = json.Unmarshal(body, &result)

	// Failed to convert the content, probably because Moodle returned an error
	if err != nil {
		err = json.Unmarshal(body, &er)
	}

	return
}
