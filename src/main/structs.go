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
	meeting array[] `json:meeting`
}
 type Meeting struct { 
	MeetingID int `json:MeetingID`
	TimeAndDate time.Time `json:TimeAndDate`
	RoomName string `json:RoomName`
	Topic string `json:Topic`
	Agenda string `json:Agenda`
	Owner string `json:Owner`
	Participants string `json:Participant`
		
 }

type Partipants struct {
	Partipants array[] `json:Participants`
}
type Participant struct {
	Name string `json:Name`
	MeetingID int `json:MeetingID`
	UserID int `json:UserID`
	OwnerID int `json:Owner`

}
