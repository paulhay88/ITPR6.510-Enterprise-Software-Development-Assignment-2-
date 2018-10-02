package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	meetings, err := meetingplannerdb.Query(`SELECT * FROM meetings`)
	check(err)

	defer meetings.Close()

	fmt.Fprintf(w, "\nMeetings: \n")

	for meetings.Next() {

		// Create new meeting. Used to store scanned query information.
		// Is appended to meetings at the end and outputed.

		var meeting Meeting

		err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
		check(err)

		// Get participants belonging to meeting

		var participants Participants

		participants, err = meetingplannerdb.Query(`SELECT * FROM participants WHERE meetingID=$1`, meeting.ID)
		check(err)

		defer participants.Close()

		fmt.Fprintf(w, "\nParticipants: \n")

		for meetings.Next() {

			// Create new meeting. Used to store scanned query information.
			// Is appended to meetings at the end and outputed.

			var meeting Meeting

			err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
			check(err)

			myMeetings.Meetings = append(myMeetings.Meetings, meeting)
		}

		myMeetings.Meetings = append(myMeetings.Meetings, meeting)
	}

	output(w, myMeetings)
}

func output(w http.ResponseWriter, myMeetings interface{}) {
	b, err := json.MarshalIndent(myMeetings, "", "\t")
	check(err)
	fmt.Fprintf(w, string(b))
}

