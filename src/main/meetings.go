package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Accepts date, room name, topic, and agenda (TODO participants)
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

	// Create meeting
	err = meetingplannerdb.QueryRow(
		`INSERT INTO meetings(dateAndTime, roomID, topic, agenda, ownerID) VALUES($1, $2, $3, $4, $5) RETURNING id`,
		meeting.TimeAndDate, meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.OwnerID).Scan(&meeting.ID)

	check(err)

	fmt.Println(meeting.ID)

	// Add user as participant
	_, err = meetingplannerdb.Exec(
		`INSERT INTO participants(meetingID, userID) VALUES($1, $2)`,
		meeting.ID, meeting.OwnerID)

	check(err)

}

// Accepts room name, topic, agenda, dateTime, (and participants TODO)
// dateTime format: "2012-09-04 14:32"
// Only updates agenda if user not owner
func updateMeeting(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var meeting = new(Meeting)
	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)
	// Decode user passed data
	err = json.NewDecoder(r.Body).Decode(&meeting)
	check(err, "Decoding error")

	fmt.Println("Datetime:")
	fmt.Println(meeting.DateTime)
	fmt.Println("Datetime end")

	// Get meeting id
	meeting.ID, err = strconv.Atoi(params.ByName("id"))
	check(err)
	// Get owner id
	err = meetingplannerdb.QueryRow(`SELECT ownerID FROM meetings WHERE id=$1`, meeting.ID).Scan(&meeting.OwnerID)
	check(err)
	// Find room associated with name
	err = meetingplannerdb.QueryRow(`SELECT id FROM rooms WHERE name=$1`, meeting.RoomName).Scan(&meeting.RoomID)
	check(err)

	// Check if user is owner
	if meeting.OwnerID != userID {

		output(w, meeting.OwnerID)
		// Not owner, only update agenda
		_, err = meetingplannerdb.Exec(
			`UPDATE meetings
			 SET 
				 agenda = $1
			 WHERE 
				  id = $2`,
			meeting.Agenda, meeting.ID)
		check(err)
	} else {

		// Owner update all that aren't null
		_, err = meetingplannerdb.Exec(
			`UPDATE meetings
			SET roomID = COALESCE(NULLIF($1, 0), roomID),
			 	topic = COALESCE(NULLIF($2, ''), topic),
				agenda = COALESCE(NULLIF($3, ''), agenda),
				dateAndTime = COALESCE(NULLIF($4, TIMESTAMP '0001-01-01 00:00:00'), dateAndTime)
			WHERE
				  id = $5`,
			meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.DateTime, meeting.ID)
		check(err)

	}
}

func deleteMeeting(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var userID int
	var meeting Meeting

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	err = meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE id=$1`, params.ByName("id")).Scan(&meeting)
	check(err)

	// Check if user is owner
	if meeting.OwnerID != userID {

		// Not owner
		return

	} else {

		// Owner can delete
		_, err := meetingplannerdb.Exec(
			`DELETE FROM participants
			 WHERE meetingID = $1`,
			meeting.ID)
		check(err)

		_, err = meetingplannerdb.Exec(
			`DELETE FROM meetings
			 WHERE id = $1`,
			meeting.ID)
		check(err)
	}

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
