package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := r.Cookie("authUser")
	check(err)
	output(w, "test")

	var se string
	err = json.NewDecoder(r.Body).Decode(&se)
	check(err)
	//se += se + ".?"
	//re := regexp.MustCompile(se)
	regularSearch := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE agenda CONTAINS $1`, se)
	fmt.Println("Agenda Return :")
	output(w, regularSearch)

}
