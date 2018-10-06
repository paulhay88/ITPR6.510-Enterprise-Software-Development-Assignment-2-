package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Outputs the saved meeting. Receives userID
func getUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting PriorMeeting

	var meeting Meeting

	// Find priorMeeting connected to userID
	priorMeetings, err := meetingplannerdb.Query(`SELECT * FROM priorMeetings WHERE userID=$1`, params.ByName("id"))
	check(err)

	defer priorMeetings.Close()

	if priorMeetings.Next() {

		err := priorMeetings.Scan(&priorMeeting.MeetingID, &priorMeeting.UserID, &priorMeeting.ID)
		check(err)

		// Find meeting connected to priorMeeting
		meetings, err := meetingplannerdb.Query(`SELECT * FROM meetings WHERE id=$1`, priorMeeting.MeetingID)
		check(err)

		defer meetings.Close()

		for meetings.Next() {
			err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
			check(err)
		}

		output(w, meeting)

	} else {
		output(w, "No settings saved.")
	}

}

func createUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting = new(PriorMeeting)

	priorMeetings, err := meetingplannerdb.Query(`SELECT * FROM priorMeetings WHERE userID=$1`, params.ByName("id"))
	check(err)

	// Check if settings already exist. Update them if they do, create them if they don't.
	if priorMeetings.Next() {
		updateUserSettings(w, r, params)

	} else {

		// Decode after potentially passing http.request to function as it will not pass correctly if its been decoded.
		err := json.NewDecoder(r.Body).Decode(&priorMeeting)
		check(err)

		_, err = meetingplannerdb.Exec(
			`INSERT INTO priorMeetings(meetingID, userID) VALUES($1, $2)`,
			priorMeeting.MeetingID, params.ByName("id"))
	}

	check(err)
}

func updateUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting = new(PriorMeeting)

	err := json.NewDecoder(r.Body).Decode(&priorMeeting)
	check(err, "error inserting")

	_, err = meetingplannerdb.Exec(
		`UPDATE priorMeetings
		 SET meetingID = $1
		 WHERE
		 	 id = $2`,
		priorMeeting.MeetingID, params.ByName("id"))
	check(err)
}

func deleteUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	_, err := meetingplannerdb.Exec(
		`DELETE FROM priorMeetings
		 WHERE userID = $1`,
		params.ByName("id"))
	check(err)
}
