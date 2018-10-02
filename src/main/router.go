package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func route() {
	router := httprouter.New()

	router.GET("/", getMeetings)
	router.POST("/meeting/create", createMeeting)
	router.PUT("/meeting/edit/:id", updateMeeting)
	router.DELETE("/meeting/delete/:id", deleteMeeting)

	// Testing
	router.POST("/outputInput", outputInput)
	log.Fatal(http.ListenAndServe(":9090", router))
}

// Turn struct
func output(w http.ResponseWriter, myStruct interface{}) {
	b, err := json.MarshalIndent(myStruct, "", "\t")
	check(err)
	fmt.Fprintf(w, string(b))
}

// Takes any input and outputs it. For testing purposes.
func outputInput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var inter interface{}

	err := json.NewDecoder(r.Body).Decode(&inter)
	check(err)

	output(w, inter)
}
