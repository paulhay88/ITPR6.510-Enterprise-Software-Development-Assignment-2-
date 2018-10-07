package main

import (
<<<<<<< HEAD
	"encoding/json"
	"net/http"
=======
>>>>>>> 2ff7dcd3488081ae2ccf620962830572fa14a3da
	"time"
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

func findMeeting(w http.ResponseWriter, r *http.Request, httprouter.Params) {
	//var findMeeting FindMeeting
	err := json.NewDecoder(r.Body).Decode(&findMeeting)
	check(err)

	/*
	// get meeting owner ID match it with user id then go threogh and find participants in other meetings and input into 
	// foundmeeting json file format for output.
	//

	// 
	//get Meeeting
		aMeeting := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE ownerID=$1`, user.UserID)
		result = foundMeeting.Scan(&id, &dateAndTime, &roomID, &topic, &agenda, &ownerID, &participants)
		if result == sql.ErrNoRows {
			output(w, "No Data :")
		} else {
			output(w, result)
		}
	// newCookie, _ := r.Cookie("authUser")

		// userName := strings.Split(newCookie.Value, ":")[0]
		// password := strings.Split(newCookie.Value, ":")[1]

		// output(w, userName)
		// output(w, password)


	*/
}

/*
	 roomID = $1,
		 topic = $2,
		 agenda = $3,
		 ownerID = $4,
		 dateAndTime = $5
*/
