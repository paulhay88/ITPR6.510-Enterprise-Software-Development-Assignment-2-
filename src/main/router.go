package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MyRouter used to call ServeHTTP
type MyRouter struct {
	meetingplannerdb *sql.DB
}

func (router *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {
			getMeetings(w, r, router.meetingplannerdb)
			return
		}

	}
	http.NotFound(w, r)
	return
}

// Any Routing functions...

func getMeetings(w http.ResponseWriter, r *http.Request, meetingplannerdb *sql.DB) {

	meetings, err := meetingplannerdb.Query(`SELECT * FROM meetings`)
	check(err)

	defer meetings.Close()

	fmt.Fprintf(w, "\nMeetings: \n")

	for meetings.Next() {
		var (
			id          int
			topic       string
			dateAndTime time.Time
			agenda      string
			roomID      int
			ownerID     int
		)

		err := meetings.Scan(&id, &topic, &dateAndTime, &agenda, &roomID, &ownerID)
		check(err)

		result := strings.Join([]string{strconv.Itoa(id), topic, dateAndTime.String(), agenda, strconv.Itoa(roomID), strconv.Itoa(ownerID)}, " ") + "\n"
		// fmt.Fprintf(w, result)
		w.Write([]byte(result))
	}
}
