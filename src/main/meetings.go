package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Outputs list of meetings
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

				err := users.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &localPass)
				check(err)

				meeting.Participants = append(meeting.Participants, user)
			}
		}
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
