package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"./moodle"
)

func main() {
	fmt.Println(moodle.NewUser("http://domain/", "username", "password", "where to save"))
}
