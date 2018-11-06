package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func route() {

	router := httprouter.New()

	mux := http.NewServeMux()

	// Seperate
	mux.Handle("/login", router)
	mux.Handle("/signup", router)
	mux.Handle("/", negroni.New(
		negroni.HandlerFunc(auth),
		negroni.Wrap(router),
	))

	// Login
	router.POST("/signup", createUser)
	router.GET("/login", loginPage)
	router.POST("/login", login)
	router.GET("/logout", logout)
	// router.GET("/login/:id", login)
	// router.POST("/createUser/", createUser)

	// Meetings
	router.GET("/allMeetings", getMeetings)
	router.GET("/", findUsersMeetings)
	router.POST("/meetings/create", createMeeting)
	router.PUT("/meetings/:id/edit", updateMeeting)
	router.DELETE("/meetings/:id/delete", deleteMeeting)

	// Query Meetings
	router.GET("/meetings", queryMeetings)

	//AgendaSearch
	router.GET("/agendaSearch", agendaSearch)

	// Rooms
	router.GET("/rooms", getRooms)
	router.POST("/rooms/create", createRoom)
	router.PUT("/rooms/:id/edit", updateRoom)
	router.DELETE("/rooms/:id/delete", deleteRoom)

	// User settings
	router.GET("/users/:id/settings", getUserSettings)
	router.POST("/users/:id/settings/create", createUserSettings)
	router.PUT("/users/:id/settings/edit", updateUserSettings)
	router.DELETE("/users/:id/settings/delete", deleteUserSettings)

	// Testing
	router.POST("/outputInput", outputInput)

	n := negroni.Classic()
	n.UseHandler(mux)
	log.Fatal(http.ListenAndServe(":9090", n))

}

// Turn struct into output
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

func auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if loggedIn(w, r, nil) {
		next(w, r)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// Single wrapper authentication
// func auth(next httprouter.Handle) httprouter.Handle {

// 	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

// 		cookie, _ := r.Cookie("authUser")

// 		userName := strings.Split(cookie.Value, ":")[0]
// 		password := strings.Split(cookie.Value, ":")[1]

// 		output(w, userName)
// 		output(w, password)

// 		next(w, r, params)
// 	}
// }
