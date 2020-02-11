package moodle

import (
	errors "../errorcodes"
)

// helperCourseListComparer compares the list of courses and returns three lists:
// 1: All the new courses
// 2: All the changed courses
// 3: All the deleted courses
// The algorithm used iterates through all elements in list1 and tries to find a
// match in list2. If an id match is found but the rest of the content is
// different it is added to the changed list and removed from list2. If it was a
// full match the element is removed from list2. If it didn't match with any
// element in list2, then it's added to the removed list. The new course list is
// list2, since the remaining courses there are all new
// Worst Case: O(n^2)
func helperCourseListComparer(list1 []Course, list2 []Course) ([]Course, []Course, []Course) {
	var modified []Course
	var deleted []Course

	for _, c := range list1 {
		var found bool = false

		for i, c1 := range list2 {
			// Check the id first
			if c.ID == c1.ID {
				found = true
				if !c.equals(c1) {
					// Modified
					modified = append(modified, c1)
				}

				// Remove from list2 (https://stackoverflow.com/a/37335777/7269000)
				list2[len(list2)-1], list2[i] = list2[i], list2[len(list2)-1]
				list2 = list2[:len(list2)-1]
				break
			}
		}

		// Deleted
		if !found {
			c.Deletedcourse = true
			deleted = append(deleted, c)
		}
	}

	// Mark the remaining courses as new
	for i, _ := range list2 {
		list2[i].Newcourse = true
	}

	return list2, modified, deleted
}

// helperUpdateToken Checks if the current moodle token is outdated, and updates
// it both in the DB and in the InfoMoodle struct. If it was successful returns
// a NoError. Otherwise it returns the corresponding error
func helperUpdateToken(moodle *InfoMoodle) errors.ErrorCode {
	resp, err := RequestAuthentication(moodle.URL, moodle.Username, moodle.Password)
	if err != nil {
		return errors.InternetError
	}

	// Check for error codes
	// Check if the request was unsuccessful
	if _, ok := resp["error"]; ok {
		return errors.ConvertMoodleError(resp["error"].(string))
	}

	if resp["token"].(string) == moodle.Token {
		// Something weird is going on
		return errors.MoodleError
	}

	// A new token exists
	moodle.Token = resp["token"].(string)
	err = SaveToken(moodle.URL, moodle.Username, moodle.Token)

	if err != nil {
		return errors.DBError
	}

	return errors.NoError
}
