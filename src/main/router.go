package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func route() {

	router := httprouter.New()
	// auth := Authenticate{router}

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
	router.POST("/outputInput", outputInput)
	log.Fatal(http.ListenAndServe(":9090", authenticate(router)))

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

// Middleware

func auth(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		cookie, _ := r.Cookie("authUser")

		userName := strings.Split(cookie.Value, ":")[0]
		password := strings.Split(cookie.Value, ":")[1]

		output(w, userName)
		output(w, password)

		next(w, r, params)
	}
}

func authenticate(router *httprouter.Router) *httprouter.Router {

	return router
}
