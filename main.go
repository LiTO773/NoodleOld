package main

import (
	"./moodle"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Register in the DB
	fmt.Println(moodle.NewUser("http://domain/", "username", "password", "where to save"))
	// fmt.Println(moodle.LoginUser("http://domain/", "username"))

	fmt.Println("===============")

	// Get the current Moodle (to be changed)
	mStruct, err := moodle.SearchMoodle("http://domain/", "username")
	fmt.Println(mStruct)
	fmt.Println(err)

	fmt.Println("===============")

	// Store the courses
	fmt.Println(moodle.CheckCourses(mStruct))
}
