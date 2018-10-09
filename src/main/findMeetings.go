package main

import (
	"database/sql"
	"fmt"
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
func findMyParticipantMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meetingCookie := r.Cokkie("authUser")
	meetings := new(Meetings)
	user := new(User)
	userName := strings.Split(meetingCookie.Value, ":")[0]
	userID := meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName)
	err := row.Scan(&user.ID)
	check(err)
	participantMeeting, err := meetingplannerdb.Query(`SELECT * FROM participants WHERE userID=$1`, user.ID)
	defer participantMeeting.Close()
	if err == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Participant of Meeting:\n")
		check(err)
		for participantMeeting.Next() {
			meeting := new(Meeting)
			p := new(Participant)
			err := participantMeeting.Scan(&p.ID, &p.MeetingID, &p.UserID)
			check(err)
			q := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE meetingID=$1`, p.MeetingID)
			err := q.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
			meetings = append(meetings, mee)
		}

	}
	output(w, meetings)
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
