package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func createMeeting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var meeting = new(Meeting)

	// User creating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&meeting.OwnerID)
	check(err)

	// Extract meeting from json
	err = json.NewDecoder(r.Body).Decode(&meeting)
	check(err)

	// Find roomID
	err = meetingplannerdb.QueryRow(`SELECT id FROM rooms WHERE name=$1`, meeting.RoomName).Scan(&meeting.RoomID)
	check(err)

	// Transaction to ensure the last meeting ID is correct
	tx, err := meetingplannerdb.Begin()
	check(err)

	// Create meeting
	err = tx.QueryRow(
		`INSERT INTO meetings(dateAndTime, roomID, topic, agenda, ownerID) VALUES($1, $2, $3, $4, $5) RETURNING id`,
		meeting.TimeAndDate, meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.OwnerID).Scan(&meeting.ID)

	check(err)

	fmt.Println(meeting.ID)

	// Find meetingID
	// err = tx.QueryRow(
	// 	`SELECT id FROM meetings ORDER BY id DESC LIMIT 1`).Scan(&meeting.ID)
	// check(err, "meetingID not found.")
	// meetingID, err := res.LastInsertId()
	// check(err)
	// meeting.ID = int(meetingID)

	// Add user as participant
	_, err = tx.Exec(
		`INSERT INTO participants(meetingID, userID) VALUES($1, $2)`,
		meeting.ID, meeting.OwnerID)

	check(err)

	err = tx.Commit()
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

func deleteMeeting(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := meetingplannerdb.Exec(
		`DELETE FROM participants
		 WHERE meetingID = $1`,
		params.ByName("id"))
	check(err)

	_, err = meetingplannerdb.Exec(
		`DELETE FROM meetings
		 WHERE id = $1`,
		params.ByName("id"))
	check(err)
}

// Outputs list of all meetings -- For testing purposes
func getMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meetingCookie, err := r.Cookie("authUser")

	output(w, meetingCookie)

	var myMeetings Meetings

	meetings, err := meetingplannerdb.Query(`SELECT * FROM meetings`)
	check(err)

	defer meetings.Close()

	for meetings.Next() {

		// Create new meeting. Used to store scanned query information.
		// Is appended to meetings at the end and outputed.

		var meeting Meeting

		err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
		check(err)

		// Get participants belonging to meeting
		participants, err := meetingplannerdb.Query(`SELECT * FROM participants WHERE meetingID=$1`, meeting.ID)
		check(err)
		defer participants.Close()

		for participants.Next() {

			var participant Participant

			err := participants.Scan(&participant.ID, &participant.MeetingID, &participant.UserID)
			check(err)

			// Get user associated with participants entry. Put this user in meetings.Participants
			users, err := meetingplannerdb.Query(`SELECT * FROM users WHERE id=$1`, participant.UserID)
			check(err)
			defer users.Close()

			for users.Next() {
				var user User
				var localPass string // keep password safe

				err := users.Scan(&user.ID, &user.UserName, &user.Name, &user.Email, &user.Phone, &localPass)
				check(err)

				meeting.Participants = append(meeting.Participants, user)
			}
		}
		myMeetings.Meetings = append(myMeetings.Meetings, meeting)
	}

	output(w, myMeetings)
}
