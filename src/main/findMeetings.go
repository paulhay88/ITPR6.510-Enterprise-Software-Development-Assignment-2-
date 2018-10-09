package main

import (
	"time"
	"regexp"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"database/sql"
	"github.com/urfave/negroni"
)

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

func findMyParticipantMeetings(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//Finds meeting by userName from Cookie
	meetingCookie := r.Cokkie("authUser")
	check(meetingCookie)
	defer meetingCookie.Close()
	userName := strings.Split(meetingCookie.Value, ":")[0]
	check(userName)
	defer userName.Close()
	//For all meetings in DB
	participantMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE participants = $1`, userName)
	check(participantMeeting)
	defer ParticipantMeeting.Close()
	if participantMeeting == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Participant of Meeting:\n")
		output(w, participantMeeting)
	}
}

func FindRoom(w http.ResponseWriter, r *http.Request, httprouter.Params){
	//need to know how to reference the info user puts in what variable to call and run through RegEx expression
	reg := regexp.MustCompile((\d+)([0-9]+))
	check(reg)
	defer reg.Close()
	roomNumber := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE RoomID = $1`, reg)  //variable based on input from user RegEx
	check(roomNumber)
	defer roomNumber.Close()
		if roomNumber == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Room Number:\n")
		output(w, roomNumber)
	}
}

func AgendaSearch(w http.ResponseWriter, r *http.Request, httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([^.?!]*(?<=[.?\s!])string(?=[\s.?!])[^.?!]*[.?!])
	check(reg)
	defer reg.Close()
	agendaReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	check(agendaReturn)
	defer agendaReturn.Close()
	if agendaReturn == sql.ErrNoRows {
		output(w, "No Data :")
	} else {
		fmt.Println("Agenda Return :\n")
		output(w, agendaReturn)
	}
}


func TopicSearch(w http.ResponseWriter, r *http.Request, s httprouter.Params){ //using s as the string to be used within the regular expression
	reg := regexp.MustCompile([a-zA-Z0-9])
	check(reg)
	defer reg.Close()
	topicReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)
	check(topicReturn)
	defer topicReturn.Close()
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