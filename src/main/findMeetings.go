package main

import (
	"time"
	"regexp"
)

func findOwnedMeeting(w http.ResponseWriter, r *http.Request, httprouter.Params) {
	var findMeeting = new(Meeting)
	defer findMeeting.Close()
	decode := json.NewDecoder(r.Body).Decode(&findMeeting)
	check(decode)
	OwnedMeeting, err := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE ownerID=$1`, user.UserID)
	check(err)
	defer OwnedMeeting.Close()
	result, err2 := foundMeeting.Scan(&id, &dateAndTime, &roomID, &topic, &agenda, &ownerID, &participants)
	defer result.Close()
	check(err2)
	if result == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Owner of :\n")
		output(w, result)
	}
}

func findMyParticipantMeetings(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//Finds meeting by userName from Cookie
	meetingCookie, err := r.Cokkie("authUser")
	check(err)
	defer meetingCookie.Close()
	userName, err2 := strings.Split(meetingCookie.Value, ":")[0]
	defer userName.Close()
	check(err2)
	//For all meetings in DB
	participantMeeting, err3 := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE participants = $1`, userName)
	defer ParticipantMeeting.Close()
	check(err3)
	if participantMeeting == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Participant of Meeting:\n")
		output(w, participantMeeting)
	}
}

func FindRoom(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//need to know how to reference the info user puts in what variable to call and run through RegEx expression
	reg, err := regexp.MustCompile((\d+)([0-9]+))
	check(err)
	defer reg.Close()
	roomNumber, err2 := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE RoomID = $1`, reg)  //variable based on input from user RegEx
	defer roomNumber.Close()
	check(err2)
	if roomNumber == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Room Number:\n")
		output(w, roomNumber)
	}
}

func AgendaSearch(w http.ResponseWriter, r *http.Request, httprouter.Params){ //using s as the string to be used within the regular expression
	reg, err := regexp.MustCompile([^.?!]*(?<=[.?\s!])string(?=[\s.?!])[^.?!]*[.?!])
	defer reg.Close()
	check(err)
	agendaReturn, err2 := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	defer agendaReturn.Close()
	check(err2)
	if agendaReturn == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Agenda Return :\n")
		output(w, agendaReturn)
	}
}


func TopicSearch(w http.ResponseWriter, r *http.Request, s httprouter.Params){ //using s as the string to be used within the regular expression
	reg, err := regexp.MustCompile([a-zA-Z0-9])
	defer reg.Close()
	check(err)
	topicReturn, err2 := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	defer topicReturn.Close()
	check(err2)
	if topicReturn == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Topic Return:\n")
		output(w, topicReturn)
	}
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