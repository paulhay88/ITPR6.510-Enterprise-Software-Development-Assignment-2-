package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Check for errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var meetingplannerdb *sql.DB

func main() {
	// defer profile.Start().Stop()

	meetingplannerdb = openDatabase()

	route()

	// Defer closing of database to end of main()
	defer meetingplannerdb.Close()

	createTables(meetingplannerdb)

	seed(meetingplannerdb)

	// Display all data from all tables
	testOverviews(meetingplannerdb)
}
