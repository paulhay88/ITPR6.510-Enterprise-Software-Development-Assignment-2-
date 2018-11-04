package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Accepts date, room name, topic, and agenda (TODO participants)
// dateTime format: "2012-09-04 14:32"
func createMeeting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var meeting = new(Meeting)

	// Owner
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
		meeting.DateTime, meeting.RoomID, meeting.Topic, meeting.Agenda, meeting.OwnerID).Scan(&meeting.ID)

	check(err)

	// Find and add participants
	for _, username := range meeting.ParticipantNames {
		var userID int

		// Find
		err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE username=$1`, username).Scan(&userID)
		check(err)

		// Cannot add self as participant
		if userID != meeting.OwnerID {

			// Add
			_, err = meetingplannerdb.Exec(
				`INSERT INTO participants(meetingID, userID) VALUES($1, $2)`,
				meeting.ID, userID)
			check(err)
		}

	}

	// Add owner as participant
	_, err = meetingplannerdb.Exec(
		`INSERT INTO participants(meetingID, userID) VALUES($1, $2)`,
		meeting.ID, meeting.OwnerID)

	check(err)

}

// Accepts room name, topic, agenda, dateTime, (and participants TODO)
// Updates only the agenda if user not owner
// Update participants:
// - to add: add new participant to the array "participants": ["test4"]
// - to remove: add existent participant to the array

func updateMeeting(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var meeting = new(Meeting)
	var userID int

	// User updating meeting
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Get user id
	err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName).Scan(&userID)
	check(err)

	// Decode user passed data
	err = json.NewDecoder(r.Body).Decode(&meeting)
	check(err, "Decoding error")

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

		// Update all that aren't null
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

		// Update participants
		for _, username := range meeting.ParticipantNames {
			var participantID int
			var errScan int

			// Find participantID
			err = meetingplannerdb.QueryRow(`SELECT id FROM users WHERE username=$1`, username).Scan(&participantID)
			check(err)

			// Find user in participants
			err = meetingplannerdb.QueryRow(`SELECT id FROM participants WHERE userID=$1 AND meetingID=$2`, participantID, meeting.ID).Scan(&errScan)

			// Check error
			if err != nil {

				// Check user not in participants
				if err == sql.ErrNoRows {

					// Cannot add self as participant
					if participantID != meeting.OwnerID {

						// Add user
						_, err = meetingplannerdb.Exec(
							`INSERT INTO participants(meetingID, userID) VALUES($1, $2)`,
							meeting.ID, participantID)
						check(err)
					}

				} else {
					check(err)
				}

				// User is already in participants
			} else {

				// Cannot remove self from participants
				if participantID != meeting.OwnerID {

					// Remove user
					_, err = meetingplannerdb.Exec(
						`DELETE FROM participants
						WHERE meetingID = $1 AND userID = $2`,
						meeting.ID, participantID)
					check(err)
				}

			}
		}

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

	}

	// Owner can delete
	_, err = meetingplannerdb.Exec(
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

		err := meetings.Scan(&meeting.ID, &meeting.Topic, &meeting.DateTime, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
		check(err)

		// Get room name
		err = meetingplannerdb.QueryRow(`SELECT name FROM rooms WHERE id=$1`, meeting.RoomID).Scan(&meeting.RoomName)
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
