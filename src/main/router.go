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
	router.GET("/", getMeetings)
	router.POST("/", createMeeting)
	router.PUT("/meeting/edit/:id", updateMeeting)

	// Testing
	router.POST("/outputInput", outputInput)
	log.Fatal(http.ListenAndServe(":9090", router))
}

// Turn struct
func output(w http.ResponseWriter, myStruct interface{}) {
	b, err := json.MarshalIndent(myStruct, "", "\t")
	check(err)
	fmt.Fprintf(w, string(b))
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

		myMeetings.Meetings = append(myMeetings.Meetings, meeting)
	}

	output(w, myMeetings)
}

func createMeeting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var meeting = new(Meeting)

	err := json.NewDecoder(r.Body).Decode(&meeting)
	check(err)

	_, err = meetingplannerdb.Exec(
		`INSERT INTO meetings(dateAndTime, roomID, topic, agenda, ownerID) VALUES($1, $2, $3, $4, $5)`,
		meeting.TimeAndDate, meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.OwnerID)

	check(err)
}

func updateMeeting(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var meeting = new(Meeting)

	err := json.NewDecoder(r.Body).Decode(&meeting)
	check(err)

	_, err = meetingplannerdb.Exec(
		`UPDATE meetings
		 SET roomID = $1, 
			 topic = $2, 
			 agenda = $3, 
			 ownerID = $4,
			 dateAndTime = $5
		 WHERE 
		 	 id = $6`,
		meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.OwnerID, meeting.TimeAndDate, params.ByName("id"))
	check(err)
}

// Takes any input and outputs it. For testing purposes.
func outputInput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var inter interface{}

	err := json.NewDecoder(r.Body).Decode(&inter)
	check(err)

	output(w, inter)
}
