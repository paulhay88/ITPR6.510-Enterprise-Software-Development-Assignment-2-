package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

/*
func findOwnedMeeting(w http.ResponseWriter, r *http.Request, httprouter.Params) {
	var findMeeting = new(Meeting)
	defer findMeeting.Close()
	decode := json.NewDecoder(r.Body).Decode(&findMeeting)
	check(decode)
	defer decode.Close()
	OwnedMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE ownerID=$1`, user.UserID)
	check(OwnedMeeting)
	defer OwnedMeeting.Close()
	result := foundMeeting.Scan(&id, &dateAndTime, &roomID, &topic, &agenda, &ownerID, &participants)
	check(result)
	defer result.Close()
	if result == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Owner of :\n")
		output(w, result)
	}
}
*/
func findUsersMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meetingCookie, err := r.Cookie("authUser")
	check(err)

	meetings := new(Meetings)
	user := new(User)

	userName := strings.Split(meetingCookie.Value, ":")[0]

	userID := meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName)
	err = userID.Scan(&user.ID)
	check(err)

	participants, err := meetingplannerdb.Query(`SELECT * FROM participants WHERE userID=$1`, user.ID)

	defer participants.Close()

	if err == sql.ErrNoRows {

		output(w, "No Data :")
		return

	} else {

		check(err)
		for participants.Next() {
			var meeting Meeting
			var p Participant

			err := participants.Scan(&p.ID, &p.MeetingID, &p.UserID)
			check(err)

			q := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE id=$1`, p.MeetingID)
			err = q.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
			meetings.Meetings = append(meetings.Meetings, meeting)
		}

	}
	output(w, "Users Meetings:\n")
	output(w, meetings.Meetings)
}

/*
func FindRoom(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//need to know how to reference the info user puts in what variable to call and run through RegEx expression
	reg := regexp.MustCompile((\d+)([0-9]+))


	roomNumber := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE RoomID = $1`, reg)  //variable based on input from user RegEx
	if roomNumber == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Room Number:\n")
		output(w, roomNumber)
	}
}

func AgendaSearch(w http.ResponseWriter, r *http.Request, httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([^.?!]*(?<=[.?\s!])string(?=[\s.?!])[^.?!]*[.?!])

	agendaReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)

	if agendaReturn == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Agenda Return :\n")
		output(w, agendaReturn)
	}
}


func TopicSearch(w http.ResponseWriter, r *http.Request, httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([a-zA-Z0-9])

	topicReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)

	if topicReturn == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Topic Return:\n")
		output(w, topicReturn)
	}
}
*/
/*
	ID           int
	TimeAndDate  time.Time
	RoomID       int
	Topic        string
	Agenda       string
	OwnerID      int
	Participants []User
*/
