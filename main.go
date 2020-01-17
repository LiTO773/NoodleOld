package main

import (
	"fmt"

	"./db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(db.GetDB())
}
