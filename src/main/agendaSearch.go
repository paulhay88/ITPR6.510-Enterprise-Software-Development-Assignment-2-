package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	var values []interface{}
	var myMeetings Meetings
	var counter int
	var searchString string
	for _, k := range []string{"sentence", "phoneNumber", "email", "keyWords", "dollar"} {
		if v, err := r.URL.Query()[k]; err {
			counter++
			var value interface{}
			if k == "sentence" {
				searchString = v[0]
				p1 := `(^\b`
				p2 := searchString
				p3 := `)?`
				p4 := `(\_[a-zA-Z0-9]+)?`
				newA := []string{p1, p2, p3, p4}
				AnotherOne := strings.Join(newA, "")
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda ~* $1", AnotherOne)
				check(err)
				for results.Next() {
					var meeting Meeting
					err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
					check(err)
					myMeetings.Meetings = append(myMeetings.Meetings, meeting)
				}
				output(w, myMeetings)
			} else if k == "phoneNumber" {
				searchString = v[1]
				p0 := `([a-zA-Z0-9\s\.\-]+)?`
				p1 := `(`
				p2 := searchString
				p3 := `)?`
				p4 := `([a-zA-Z0-9\s\.\-]+)?`
				newA := []string{p0, p1, p2, p3, p4}
				AnotherOne := strings.Join(newA, "")
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda ~* $1", AnotherOne)
				check(err)
				for results.Next() {
					var meeting Meeting
					err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
					check(err)
					myMeetings.Meetings = append(myMeetings.Meetings, meeting)
				}
				output(w, myMeetings)
			} else if k == "email" {
				searchString = v[1]
				p0 := `([a-zA-Z0-9\s\-\.]+)(\s)`
				p1 := `(`
				p2 := searchString
				p3 := `)?`
				p4 := `(\@)?([a-zA-Z]+)?(\.)?(com|au|co.nz)?([a-zA-Z0-9\s\-\.]+)`
				newA := []string{p0, p1, p2, p3, p4}
				AnotherOne := strings.Join(newA, "")
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda ~* $1", AnotherOne)
				check(err)
				for results.Next() {
					var meeting Meeting
					err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
					check(err)
					myMeetings.Meetings = append(myMeetings.Meetings, meeting)
				}
				output(w, myMeetings)
				//Working uptill here
			} else if k == "keyWords" {
				k1 := v[0]
				k2 := v[1]
				k3 := v[2]
				p0 := `(`
				p1 := k1
				p2 := `)([a-zA-Z\n\s\.\w]+)?(`
				p3 := k2
				p4 := `)([a-zA-Z\n\s\.\w]+)?(`
				p5 := k3
				p6 := `)`
				p7 := `|`
				newA := []string{p0, p1, p2, p3, p4, p5, p6, p7, p0, p3, p2, p5, p4, p1, p6, p7, p0, p5, p2, p1, p4, p3, p6}
				AnotherOne := strings.Join(newA, "")
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda ~* $1", AnotherOne)
				check(err)
				for results.Next() {
					var meeting Meeting
					err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
					check(err)
					myMeetings.Meetings = append(myMeetings.Meetings, meeting)
				}
				output(w, myMeetings)
			} else if k == "dollar" {
				searchString = v[1]
				p1 := `([a-zA-z0-9\n\.\s\-])+(\\`
				p2 := searchString
				p3 := `\b)([a-zA-z0-9\n\.\s\-])+`

				newA := []string{p1, p2, p3}
				fmt.Println(newA)
				AnotherOne := strings.Join(newA, "")
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda ~* $1", AnotherOne)
				check(err)
				for results.Next() {
					var meeting Meeting
					err := results.Scan(&meeting.ID, &meeting.Topic, &meeting.TimeAndDate, &meeting.Agenda, &meeting.RoomID, &meeting.OwnerID)
					check(err)
					myMeetings.Meetings = append(myMeetings.Meetings, meeting)
				}
				output(w, myMeetings)
				//regExpression = `([a-zA-Z0-9\ \-\.\n]+)(\ )(\$[0-9\.]+)(\ )([a-zA-Z\ \-\.\n]+)`
			} else {
				value = v[0]
			}
			values = append(values, value)

		}
		if len(values) == 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

	}

}
