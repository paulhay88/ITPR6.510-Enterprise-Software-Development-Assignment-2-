package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func findUsersMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meetingCookie, err := r.Cookie("authUser")
	check(err)

	output(w, "test")

	meetings := new(Meetings)
	user := new(User)

	userName := strings.Split(meetingCookie.Value, ":")[0]

	userID := meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName)
	err = userID.Scan(&user.ID)
	check(err)

	participants, err := meetingplannerdb.Query(`SELECT * FROM participants WHERE userID=$1`, user.ID)

	defer participants.Close()

	if err == sql.ErrNoRows {

		output(w, "No Data")
		return
	}

	check(err)
	for participants.Next() {
		var meeting Meeting
		var p Participant

		err := participants.Scan(&p.ID, &p.MeetingID, &p.UserID)
		check(err)

		q := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE id=$1`, p.MeetingID)
		err = q.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
		meetings.Meetings = append(meetings.Meetings, meeting)

	}
	output(w, meetings.Meetings)
}

func queryMeetings(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var meetings Meetings

	var values []interface{}
	var where []string
	var counter int

	for _, k := range []string{"dateAndTime", "topic", "roomName", "ownerName"} {
		if v, err := r.URL.Query()[k]; err {
			counter++
			var value interface{}

			if k == "roomName" {
				fmt.Println(k)
				_ = meetingplannerdb.QueryRow("SELECT id FROM rooms WHERE name=$1", v[0]).Scan(&value)
				k = "roomID"
			} else if k == "ownerName" {
				_ = meetingplannerdb.QueryRow("SELECT id FROM users WHERE name=$1", v[0]).Scan(&value)
				k = "ownerID"
			} else {
				value = v[0]
			}
			values = append(values, value)
			where = append(where, fmt.Sprintf("%s = $%d", k, counter))
		}
	}

	// Return to home if query cleared.
	if len(values) == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE "+strings.Join(where, " AND "), values...)
	check(err)

	for results.Next() {
		var meeting Meeting
		err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
		check(err)

		meetings.Meetings = append(meetings.Meetings, meeting)
	}

	output(w, meetings)
}

// func AgendaSearch(w http.ResponseWriter, r *http.Request, httprouter.Params){ //using s as the string to be used within the regular expression
// 	reg := regexp.MustCompile([^.?!]*(?<=[.?\s!])string(?=[\s.?!])[^.?!]*[.?!])

// 	agendaReturn := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE Agenda = $1`, reg)

// 	if agendaReturn == sql.ErrNoRows {
// 		output(w, "No Data :")
// 	} else {
// 		fmt.Println("Agenda Return :\n")
// 		output(w, agendaReturn)
// 	}
// }

// -------------------------------------------
// r.URL.Query() functionality as passed parameter.
// Built before discovering r.URL.Query()
// -------------------------------------------
// var queryExp = regexp.MustCompile(`(\?)?(?P<Topic>topic+)?(\=+)?(?P<TopicResult>[a-zA-z\_]+)?(\&+)?(?P<DateAndTime>dateAndTime)?(\=+)?(?P<DateAndTimeResult>[a-zA-Z0-9\.\-\\\:\/]+)?(\&)?(?P<RoomNAme>roomName)?(\=)?(?P<RoomResults>[a-zA-Z0-9\-\\\/]+)?(\&)?(?P<OwnerName>ownerName)?(\=)?(?P<OwnerResults>[a-zA-Z0-9\-\\\/\.]+)?`)

// 	queryString := queryExp.FindStringSubmatch(params.ByName("query"))

// 	subGroups := make(map[string]string)
// 	for i, name := range queryExp.SubexpNames() {
// 		if i != 0 && name != "" {
// 			subGroups[name] = queryString[i]
// 		}
// 	}
// 	output(w, subGroups["TopicResult"])
