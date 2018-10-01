package main

import "time"

// Users array
type Users struct {
	Users []User `json:"users"`
}

// User details
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Rooms array
type Rooms struct {
	Rooms []Room `json:"rooms"`
}

// Room details
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PriorMeetings array
type PriorMeetings struct {
	PriorMeetings []PriorMeeting `json:"priorMeetings"`
}

// PriorMeeting references
type PriorMeeting struct {
	ID        int `json:"id"`
	MeetingID int `json:"meetingID"`
	UserID    int `json:"userID"`
}

type Meetings struct {
	Meetings []Meeting `json:"meeting"`
}
type Meeting struct {
	ID          int       `json:"id"`
	TimeAndDate time.Time `json:"timeAndDate"`
	RoomID      int       `json:"roomID"`
	Topic       string    `json:"topic"`
	Agenda      string    `json:"agenda"`
	OwnerID     int       `json:"ownerID"`
}

type Partipants struct {
	Partipants []Participant `json:"participants"`
}
type Participant struct {
	Name      string `json:"name"`
	MeetingID int    `json:"meetingID"`
	UserID    int    `json:"userID"`
	OwnerID   int    `json:"owner"`
}
