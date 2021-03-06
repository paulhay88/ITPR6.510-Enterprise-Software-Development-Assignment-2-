package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Outputs the saved meeting. Receives userID
func getUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting PriorMeeting

	var meeting Meeting

	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Get user id
	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	// Find priorMeeting connected to userID
	priorMeetings, err := meetingplannerdb.Query(`SELECT * FROM priorMeetings WHERE userID=$1`, userID)
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
			err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.DateTime, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
			check(err)
		}

		output(w, meeting)

	} else {
		output(w, "No settings saved.")
	}

}

func createUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting = new(PriorMeeting)

	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Get user id
	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	priorMeetings, err := meetingplannerdb.Query(`SELECT * FROM priorMeetings WHERE userID=$1`, userID)
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
			priorMeeting.MeetingID, userID)
	}

	check(err)
}

func updateUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var priorMeeting = new(PriorMeeting)

	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Get user id
	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	err = json.NewDecoder(r.Body).Decode(&priorMeeting)
	check(err, "error inserting")

	_, err = meetingplannerdb.Exec(
		`UPDATE priorMeetings
		 SET meetingID = $1
		 WHERE
		 	 id = $2`,
		priorMeeting.MeetingID, userID)
	check(err)
}

func deleteUserSettings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Get user id
	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	_, err = meetingplannerdb.Exec(
		`DELETE FROM priorMeetings
		 WHERE userID = $1`,
		userID)
	check(err)
}
