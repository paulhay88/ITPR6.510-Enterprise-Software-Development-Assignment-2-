package main

// Users array
type Users struct {
	Users []User `json:"users"`
}

// User details
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
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
	Meetings []Meeting `json:"meetings"`
}
type Meeting struct {
	ID       int    `json:"id"`
	RoomID   int    `json:"roomID"`
	Topic    string `json:"topic"`
	Agenda   string `json:"agenda"`
	OwnerID  int    `json:"ownerID"`
	DateTime string `json:"dateTime"`
	// For decoding
	ParticipantNames []string `json:"participants"`
	RoomName         string   `json:"roomName"`

	// For conveniance storage
	Participants []User
}

type Partipants struct {
	Partipants []Participant `json:"participants"`
}
type Participant struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	MeetingID int    `json:"meetingID"`
	UserID    int    `json:"userID"`
	OwnerID   int    `json:"ownerID"`
}

type agendaSearchStruct struct {
	KeyWord string `json:"keyWord"`
}
