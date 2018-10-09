package main

import (
	"time"
	"regexp"
)

func findOwnedMeeting(w http.ResponseWriter, r *http.Request, httprouter.Params) {
	var findMeeting = new(Meeting)
	err := json.NewDecoder(r.Body).Decode(&findMeeting)
	check(err)
	OwnedMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE ownerID=$1`, user.UserID)
	result := foundMeeting.Scan(&id, &dateAndTime, &roomID, &topic, &agenda, &ownerID, &participants)
	if result == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Owner of :\n")
		output(w, result)
	}
}

func findMyParticipantMeetings(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//Finds meeting by userName from Cookie
	meetingCookie, _ := r.Cokkie("authUser")
	userName, err := strings.Split(meetingCookie.Value, ":")[0]
	check(err)
	//For all meetings in DB
	ParticipantMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE participants = $1`, userName)
	if result == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Owner of :\n")
		output(w, ParticipantMeeting)
	}
}

func FindRoom(w http.ResponseWriter, r *http.Request, s httprouter.Params){
	//need to know how to reference the info user puts in what variable to call and run through RegEx expression
	roomNumber := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE RoomID = $1`, s)  //variable based on input from user RegEx
	output(w, roomNumber)
}

func AgendaSearch(w http.ResponseWriter, r *http.Request, s httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([^.?!]*(?<=[.?\s!])string(?=[\s.?!])[^.?!]*[.?!])
	agendaReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	output(w, agendaReturn)
}


func TopicSearch(w http.ResponseWriter, r *http.Request, s httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([a-zA-Z0-9])
	topicReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	output(w, topicReturn)
}

/*
	ID           int       
	TimeAndDate  time.Time 
	RoomID       int      
	Topic        string    
	Agenda       string    
	OwnerID      int       
	Participants []User
*/