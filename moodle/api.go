package moodle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// RequestAuthentication requests Moodle for a user's token
// Resquest: hostname/login/token.php?service=moodle_mobile_app
// Body: username: username, password: password
// Returns the user's token in the Moodle or an error code if Moodle Web Services aren't allowed in the server
func RequestAuthentication(hostname string, username string, password string) (map[string]interface{}, error) {
	var result map[string]interface{}

	resp, err := http.PostForm(hostname+"/login/token.php?service=moodle_mobile_app", url.Values{
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
	json.Unmarshal([]byte(body), &result)

	return result, nil
}
