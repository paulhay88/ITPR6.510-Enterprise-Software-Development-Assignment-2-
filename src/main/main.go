package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Check for errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// defer profile.Start().Stop()

	meetingplannerdb := openDatabase()

	// Defer closing of database to end of main()
	defer meetingplannerdb.Close()

	createTables(meetingplannerdb)

	overview, err := meetingplannerdb.Query(`SELECT * FROM users`)
	check(err)

	defer overview.Close()

	for overview.Next() {
		var (
			name     string
			phone    string
			email    string
			password string
		)

		err := overview.Scan(&name, &phone, &email, &password)
		check(err)

		fmt.Println(name, phone, email, password)
	}

	fmt.Println(meetingplannerdb)
}
