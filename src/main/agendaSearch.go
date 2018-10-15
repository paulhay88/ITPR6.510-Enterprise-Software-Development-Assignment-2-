package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	output(w, "test")
	_, err := r.Cookie("authUser")
	check(err)
	output(w, "test")
	var se agendaSearchStruct
	//decode into a struct format to handle better possibly ? ? ?

	err = json.NewDecoder(r.Body).Decode(&se)
	check(err)
	//se += se + ".?"
	//re := regexp.MustCompile(se)
	regularSearch := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE agenda CONTAINS $1`, se)
	fmt.Println("Agenda Return :")
	output(w, regularSearch)

}
