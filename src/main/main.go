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
	router := &MyRouter{}
	http.ListenAndServe(":9090", router)

	meetingplannerdb := openDatabase()

	// Defer closing of database to end of main()
	defer meetingplannerdb.Close()

	createTables(meetingplannerdb)

	seed(meetingplannerdb)

	// Display all data from all tables
	testOverviews(meetingplannerdb)
	// fmt.Println(meetingplannerdb)
}
