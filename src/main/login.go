package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type LogInUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
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

func validateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var logInUser LogInUser

	err := json.NewDecoder(r.Body).Decode(&logInUser)
	check(err)

	// Get user with the POSTed details
	userRow := meetingplannerdb.QueryRow(`SELECT * FROM users WHERE userName=$1 AND password=$2`, logInUser.UserName, logInUser.Password)

	var user User

	err = userRow.Scan(&user.ID, &user.UserName, &user.Name, &user.Phone, &user.Email, &user.Password)

	// Check if user exists
	if err == sql.ErrNoRows {
		http.Redirect(w, r, "/signup", http.StatusUnauthorized)
		output(w, "Computer says 'No'.")
	} else {
		expiration := time.Now().Add(1 * 24 * time.Hour)
		cookie := http.Cookie{Name: user.UserName, Value: user.Password, Expires: expiration}
		http.SetCookie(w, &cookie)

		for _, cookie := range r.Cookies() {

			output(w, cookie.Name)
			output(w, cookie.Value)
		}
	}

}
