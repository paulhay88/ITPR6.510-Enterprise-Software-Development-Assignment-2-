package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	//var meetings Meetings
	var values []interface{}

	var counter int
	agendaCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(agendaCookie.Value, ":")[0]
	userID := meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName)

	for _, k := range []string{"sentence", "phoneNumbe", "email", "keyWords", "dollar"} {
		if v, err := r.URL.Query()[k]; err {
			counter++
			var value interface{}

			if k == "sentence" {
				output(w, v[0])
				//x = v[0]
				//y := meetingplannerdb.Query("SELECT regexp_matches() FROM meetings WHERE id=$1", userID)
			} else if k == "phoneNumbe" {
				output(w, v[0])
			} else if k == "email" {
				output(w, v[0])
			} else if k == "keyWords" {
				output(w, v[0])
			} else if k == "dollar" {
				output(w, v[0])
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
	output(w, userID)
	//output(w, r.URL.Query())
	/*
		agendaCookie, err := r.Cookie("authUser")
		check(err)
		//getting cokkie
		output(w, "test1")

		var se = new(agendaSearchStruct)
		//getting username from cookie
		userName := strings.Split(agendaCookie.Value, ":")[0]
		//decoding the body aka search field
		err = json.NewDecoder(r.Body).Decode(&se)
		check(err)
		//full query using contians
		searchQuery := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE userName=$1 AND agenda CONTAINS=$2`, userName, se)
		output(w, searchQuery)
	*/
}
