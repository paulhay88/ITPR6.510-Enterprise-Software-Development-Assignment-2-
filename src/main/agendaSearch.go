package main

import (
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
			} else if k == "keyWords" {
				output(w, v[0])
				/*
					for len(v){
						insert into array of strings
						concatanate those and search for them
					}
						var var1 = "string"
						var var2 = "string"
						var var3 = "string"
				*/
				//regExpression = `([a-zA-Z\ \-\.\n]+)(\ )( ` + var1 + `)(\ )([a-zA-Z\ \-\.\n]+)(\ )(` + var2 + `)(\ )([a-zA-Z\ \-\.\n]+)(\ )(` + var3 + `)(\ )([a-zA-Z\ \-\.\n]+)`
			} else if k == "dollar" {
				output(w, v[0])
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
