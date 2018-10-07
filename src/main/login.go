package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type LogInUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Is user logged in
func loggedIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) bool {

	cookie, err := r.Cookie("authUser")

	if err != nil {
		return false
	}

	userName := strings.Split(cookie.Value, ":")[0]
	password := strings.Split(cookie.Value, ":")[1]

	userRow := meetingplannerdb.QueryRow(`SELECT * FROM users WHERE userName=$1 AND password=$2`, userName, password)

	var user User

	err = userRow.Scan(&user.ID, &user.UserName, &user.Name, &user.Phone, &user.Email, &user.Password)

	// Check if user exists
	if err != sql.ErrNoRows {
		return true
	}

	return false
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user = new(User)
	err := json.NewDecoder(r.Body).Decode(&user)
	check(err)

	userRow := meetingplannerdb.QueryRow(`SELECT * FROM users WHERE userName=$1`, user.UserName)

	err = userRow.Scan(&user.ID, &user.UserName, &user.Name, &user.Phone, &user.Email, &user.Password)

	// Check if username already exists
	if err == sql.ErrNoRows {
		http.Redirect(w, r, "/signup", http.StatusUnauthorized)
		output(w, "Username already exists, please choose a different one.")
	} else {

		// Create new user
		_, err = meetingplannerdb.Exec(
			`INSERT INTO users(name, userName, phone, email, password) VALUES($1, $2, $3, $4, $5)`,
			user.Name, user.UserName, user.Phone, user.Email, user.Password)
		check(err)
	}

}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var logInUser LogInUser

	if loggedIn(w, r, nil) {
		output(w, "User is already logged in.")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&logInUser)
	check(err)

	// Get user with the POSTed details
	userRow := meetingplannerdb.QueryRow(`SELECT * FROM users WHERE userName=$1 AND password=$2`, logInUser.UserName, logInUser.Password)

	var user User

	err = userRow.Scan(&user.ID, &user.UserName, &user.Name, &user.Phone, &user.Email, &user.Password)

	// Check if user exists
	if err == sql.ErrNoRows {
		output(w, "Your username or password was incorrect. Please try again.")
	} else {
		expiration := time.Now().Add(1 * 24 * time.Hour)
		cookie := http.Cookie{Name: "authUser", Value: user.UserName + ":" + user.Password, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if !loggedIn(w, r, nil) {
		output(w, "User is already logged out.")
		return
	}

	cookie := http.Cookie{Name: "authUser", Value: ":"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func loginPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	output(w, "Please login.")
}
