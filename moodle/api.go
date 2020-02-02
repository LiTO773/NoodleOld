package moodle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
		log.Fatalln(err)
		return result, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return result, err
	}

	// Get the content
	err = json.Unmarshal([]byte(body), &result)

	return result, err
}

// GetSiteInfo requests the information about the user and the moodle
// It can also be used to test if the token is still valid
// Request: <host>/webservice/rest/server.php?moodlewsrestformat=json
// Body: wstoken: token, wsfunction: core_webservice_get_site_info
// Returns a MoodleInfo object if the request was successful (even if an error
// was received) or an error if it wasn't able to connect to the server
func GetSiteInfo(host string, token string) (MoodleInfo, error) {
	var result MoodleInfo

	resp, err := http.PostForm(host+"webservice/rest/server.php?moodlewsrestformat=json", url.Values{
		"wstoken":    {token},
		"wsfunction": {"core_webservice_get_site_info"},
	})
	if err != nil {
		log.Fatalln(err)
		return result, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return result, err
	}

	// Get the content
	err = json.Unmarshal([]byte(body), &result)

	return result, err
}
