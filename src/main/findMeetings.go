package main

import (
	"os/user"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
type FindMeeting struct {
	ID           int       `json:"id"`
	TimeAndDate  time.Time `json:"timeAndDate"`
	RoomID       int       `json:"roomID"`
	Topic        string    `json:"topic"`
	Agenda       string    `json:"agenda"`
	OwnerID      int       `json:"ownerID"`
	Participants []User
}

func findMeeting(w http.ResponseWriter, r *http.Request) {
	var findMeeting FindMeeting
	err := json.NewDecoder(r.Body).Decode(&findMeeting)
	check(err)
	//get cookie
	http.GetCookie(r, &cookie)
	for _, cookie := range r.Cookies() {
		if cookie.Name == user.Name{
		output(w, cookie.Name)
		//get Meeeting 
		aMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE ownerID=$1`, user.UserID)
		result = foundMeeting.Scan(&id, &dateAndTime, &roomID, &topic, &agenda, &ownerID, &participants)
		if result == sql.ErrNoRows {
			output(w, "No Data :")
		} else {
			output(w, result)
		}

	}
}




/*
		 roomID = $1, 
			 topic = $2, 
			 agenda = $3, 
			 ownerID = $4,
			 dateAndTime = $5
*/