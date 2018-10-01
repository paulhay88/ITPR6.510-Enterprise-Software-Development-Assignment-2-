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
	meetings []Meeting
}
 type Meeting struct { 
	MeetingID int
	TimeAndDate time.Time
	RoomName string
	Topic string
	Agenda string
	Owner string
	Participants string
		
 }

type Partipants struct {
	Partipants array[]
}
type Participant struct {
	Name string
	MeetingID int
	UserID int
	OwnerID int
}
