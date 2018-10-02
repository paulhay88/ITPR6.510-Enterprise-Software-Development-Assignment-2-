package	main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
	type LogInUser struct {
		userName string 
		password string
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
				//validate
			}
		}
	}