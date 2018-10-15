package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
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

}
