package	main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
type LogInUser struct {
	userName string 
	password string
}
func routeLog() {
	router := httprouter.New()
	router.GET("/login/", validateUser)
	router.POST("/", createUser)
	log.Fatal(http.ListenAndServe(":9090", router))
}
	func createUser(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
		var user = new(User)
		err := json.NewDecoder(r.Body).Decode(&user)
	check(err)

	_, err = users.Exec(
		`INSERT INTO users(name, phone, email, password) VALUES($1, $2, $3, $4)`,
		users.name, users.phone, users.email, users.password)
	check(err)

	}
	

func validateUser(w http.ResponseWriter, r *http.Request,  _ httprouter.Params){

		var checkUser Users
		var logInUser logInUser
		users, err := users.Query(`SELECT * FROM users`)
		check(err)
		defer users.Close
		err := json.NewDecoder(r.Body).Decode(&user)
		for users.Next(){
			var user User
			err := users.Scan(&name, &phone, &email, &password)
			check(err)
			//coompare 
			if user.Name == logInUser.userName && user.Password == logInUser.password{
				/*
				validatation here and proceed to create cookie
				setting a cookie:
				*/
				
				expiration := time.Now().Add(365 * 24 * time.Hour)
				cookie := http.Cookie{Name: user.Name, Value: user.Password, Expires: expiration, Secure: true}
				http.SetCookie(w, &cookie)
				for	_, cookie := range r.Cookies(){  //new variable should change ? 

					fmt.Fprint(w, cookie.Name)
					} //should return cookies 
				message = "User Accepted"
				http.Redirect(w, r, "/", http.StatusFound)
			}
			else{
				message = "User Not accepted"
				http.Redirect(w, r, "/login/", http.StatusFound)
			}
		}
	}

	//COOKIES
	/*
	http.SetCookie(w ResponseWriter, cookie *Cookie)
	type Cookie struct {
	    Name       string
	    Value      string
	    Path       string
	    Domain     string
	    Expires    time.Time
	    RawExpires string
	    MaxAge   int
	    Secure   bool
	    HttpOnly bool
	    Raw      string
	    Unparsed []string // Raw text of unparsed attribute-value pairs
	}
	*/
//get a cookie that has been set
//cookie, _ := r.Cookie("username")
//fmt.Fprint(w, cookie)
	//Here is another way to get a cookie
	/*
	for _, cookie := range r.Cookies() {
	    fmt.Fprint(w, cookie.Name)
	}
	*/
