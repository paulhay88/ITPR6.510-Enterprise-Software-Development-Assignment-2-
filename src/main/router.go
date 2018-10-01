package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func route() {
	router := httprouter.New()
	// router.Handle("/", getHandler).Methods("GET")
	// router.Handle("/", postHandler).Methods("POST")
	// router.Handle("/", puthandler).Methods("PUT")
	// router.handle("/", deletehandle).Methods("DELETE")
	router.GET("/", getMeetings)
	log.Fatal(http.ListenAndServe(":9090", router))
}

// Any Routing functions...

func getMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var myMeetings Meetings

	// var myMeetings = Meetings{[]Meeting{Meeting{1, time.Now(), 1, "topic", "agenda", 1}}}

	// var myMeetings = Meeting{1, time.Now(), 1, "topic", "agenda", 1}

	// b, err := json.MarshalIndent(myMeetings, "", "\t")
	// check(err)
	// fmt.Fprintf(w, string(b))

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

		// err := meetings.Scan(&id, &topic, &dateAndTime, &agenda, &roomID, &ownerID)
		// check(err)

		err := meetings.Scan(&id, &topic, &dateAndTime, &agenda, &roomID, &ownerID)
		check(err)

		// result := strings.Join([]string{strconv.Itoa(id), topic, dateAndTime.String(), agenda, strconv.Itoa(roomID), strconv.Itoa(ownerID)}, " ") + "\n"
		// fmt.Fprintf(w, result)
		// w.Write([]byte(result))
	}
}
