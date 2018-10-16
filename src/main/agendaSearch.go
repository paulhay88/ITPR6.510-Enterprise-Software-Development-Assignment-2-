package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func agendaSearch(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	//var meetings Meetings
	var values []interface{}

	var counter int
	//agendaCookie, err := r.Cookie("authUser")
	//check(err)
	//userName := strings.Split(agendaCookie.Value, ":")[0]
	//userID := meetingplannerdb.QueryRow(`SELECT id FROM users WHERE userName=$1`, userName)
	//var regExpression string
	var searchString string
	for _, k := range []string{"sentence", "phoneNumbe", "email", "keyWords", "dollar"} {
		if v, err := r.URL.Query()[k]; err {
			counter++
			var value interface{}

			if k == "sentence" {
				searchString = v[0]
				// Creating the RegEx as a String using []strings to build full query

				p1 := `(^\b)(`
				// It seems to be the output function that doesn't work on these kinds of strings. Use this instead:
				fmt.Fprintf(w, p1)

				p2 := searchString
				p3 := `)(\ )([a-zA-Z0-9\ \-]+)?(\.)?`
				newA := []string{p1, p2, p3}
				AnotherOne := strings.Join(newA, "")
				output(w, AnotherOne) //Doesn't Work

				regExToSearch := []string{`(^\b)(`, searchString, `)(\ )([a-zA-Z0-9\ \-]+)?(\.)?`}
				anotherString := strings.Join(regExToSearch, "")
				YetAnotherString, _ := regexp.Compile(anotherString)
				output(w, YetAnotherString) //This Doesn't

				Comp, _ := regexp.Compile("[\\D]") //This Works
				y := Comp.FindString("T")
				output(w, y)

				//output(w, anotherString)
				convertedString := regexp.MustCompile(anotherString)
				//convertedString.FindAll()
				output(w, convertedString)

				//results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda LIKE $1", convertedString)
				//check(err)
				//output(w, results)

			} else if k == "phoneNumbe" {
				output(w, v[0])
				//regExpression = `([a-zA-Z\ \-\$\.]+)?(\ )?([0-9\ \-\.]+)([a-zA-Z\ \-\.]+)`
			} else if k == "email" {
				output(w, v[0])
				//regExpression = `([a-zA-Z\ \-]+)(\ )([a-zA-Z0-9\@]+)(\.)(com|net|us|ca|org|au|co\.nz)([a-zA-Z\ \-\.]+)`
			} else if k == "keyWords" {
				output(w, v[0])
				/*
					var var1 = "string"
					var var2 = "string"
					var var3 = "string"
				*/
				//regExpression = `([a-zA-Z\ \-\.\n]+)(\ )( ` + var1 + `)(\ )([a-zA-Z\ \-\.\n]+)(\ )(` + var2 + `)(\ )([a-zA-Z\ \-\.\n]+)(\ )(` + var3 + `)(\ )([a-zA-Z\ \-\.\n]+)`
			} else if k == "dollar" {
				output(w, v[0])
				//regExpression = `([a-zA-Z0-9\ \-\.\n]+)(\ )(\$[0-9\.]+)(\ )([a-zA-Z\ \-\.\n]+)`
			} else {
				value = v[0]
			}
			values = append(values, value)

		}
		if len(values) == 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

	}

}
