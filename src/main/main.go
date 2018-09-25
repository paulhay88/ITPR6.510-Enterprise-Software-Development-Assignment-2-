package main

import (
	"log"
	"net/http"

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

	router := &MyRouter{meetingplannerdb}
	http.ListenAndServe(":9090", router)

	// Defer closing of database to end of main()
	defer meetingplannerdb.Close()

	createTables(meetingplannerdb)

	seed(meetingplannerdb)

	// Display all data from all tables
	testOverviews(meetingplannerdb)
}
