package main

// Users array
type Users struct {
	users []User `json:"users"`
}

// User details
type User struct {
	id       int    `json:"id`
	name     string `json:"name"`
	phone    string `json:"phone"`
	email    string `json:"email"`
	password string `json:"password"`
}

// Rooms array
type Rooms struct {
	rooms []Room `json:"rooms"`
}

// Room details
type Room struct {
	id   int    `json:"id`
	name string `json:"name"`
}

// PriorMeetings array
type PriorMeetings struct {
	priorMeetings []PriorMeeting `json:"priorMeetings"`
}

// PriorMeeting references
type PriorMeeting struct {
	id        int `json:"id`
	meetingID int `json:"meetinID`
	userID    int `json:"userID`
}


type Meetings struct {
	meeting array[] `json:"meeting"`
}
 type Meeting struct { 
	meetingID int `json:"meetingID"`
	timeAndDate time.Time `json:"timeAndDate"`
	roomName string `json:"roomName"`
	topic string `json:"topic"`
	agenda string `json:"agenda"`
	owner string `json:"owner"`
	participants string `json:"participant"`
		
 }

type Partipants struct {
	Partipants array[] `json:"Participants"`
}
type Participant struct {
	name string `json:"name"`
	meetingID int `json:"meetingID"`
	userID int `json:"userID"`
	ownerID int `json:"owner"`

}
