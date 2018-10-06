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
<<<<<<< HEAD

	router.PUT("/meeting/edit/:id", updateMeeting)
	router.DELETE("/meeting/delete/:id", deleteMeeting)
=======
	router.PUT("/meeting/:id/edit", updateMeeting)
	router.DELETE("/meeting/:id/delete", deleteMeeting)

	// Users

	// User settings
	router.GET("/users/:id/settings", getUserSettings)
	router.POST("/users/:id/settings/create", createUserSettings)
	router.PUT("/users/:id/settings/edit", updateUserSettings)
	router.DELETE("/users/:id/settings/delete", deleteUserSettings)
>>>>>>> 9ea6d639422ec33fd2f95961ccf4a65d4763fa50

	// Testing
	router.POST("/outputInput", outputInput)
	log.Fatal(http.ListenAndServe(":9090", router))

	// Login
	router.GET("/login/", _)
	router.GET("/login/:id", login)
	router.POST("/createUser/", createUser)

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
