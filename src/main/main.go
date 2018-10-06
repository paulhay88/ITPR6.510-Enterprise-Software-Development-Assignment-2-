package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Check for errors
func check(err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			fmt.Println(message[0])
		}
		log.Fatal(err)
	}
}

var meetingplannerdb *sql.DB

func main() {
	// defer profile.Start().Stop()

	meetingplannerdb = openDatabase()

	// Defer closing of database to end of main()
	defer meetingplannerdb.Close()

	createTables(meetingplannerdb)

	seed(meetingplannerdb)

	route()

	// routeLog()

	// Display all data from all tables
	testOverviews(meetingplannerdb)
}
