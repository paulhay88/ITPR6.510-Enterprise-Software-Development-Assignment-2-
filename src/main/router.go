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

	// Meetings
	router.GET("/", getMeetings)
	router.POST("/meeting/create", createMeeting)
	router.PUT("/meeting/:id/edit", updateMeeting)
	router.DELETE("/meeting/:id/delete", deleteMeeting)

	// Users

	// User settings
	router.GET("/users/:id/settings", getUserSettings)
	router.POST("/users/:id/settings/create", createUserSettings)
	router.PUT("/users/:id/settings/edit", updateUserSettings)
	router.DELETE("/users/:id/settings/delete", deleteUserSettings)

	// Login
	router.POST("/signup", createUser)
	router.POST("/login", validateUser)
	// router.GET("/login/:id", login)
	// router.POST("/createUser/", createUser)

	// Testing
	router.GET("/getCookies", getCookies)
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
