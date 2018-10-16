package main

import (
	"net/http"
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
				output(w, searchString)
				//regExToSearch := `(^\\` + `b)(` + searchString + `)(\\` + ` )(a-zA-Z0-9\\` + ` \\` + `-]+?(\\` + `.)?`
				/*
					Top one tries to break the string up with double \\ but returns \\\\
					----------- switch these two out-------------
					Bottom One returns double \\
				*/
				sillVar := `"\"`
				// ------ I even tried setting it as its own charecter
				regExToSearch := []string{`(^\b)(`, searchString, `)(\ )([a-zA-Z0-9\ \-]+)?(\.)?`}
				convertedString := strings.Join(regExToSearch, "")
				output(w, convertedString)
				output(w, sillVar)
				results, err := meetingplannerdb.Query("SELECT * FROM meetings WHERE agenda LIKE $1", convertedString)
				check(err)
				output(w, results)
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

	//results, err := meetingplannerdb.Query("SELECT "+strings.Join(regExpression, " FROM meetings WHERE id=$1"), userID)

	//output(w, userID)

	//output(w, r.URL.Query())
	/*
		agendaCookie, err := r.Cookie("authUser")
		check(err)
		//getting cokkie
		output(w, "test1")

		var se = new(agendaSearchStruct)
		//getting username from cookie
		userName := strings.Split(agendaCookie.Value, ":")[0]
		//decoding the body aka search field
		err = json.NewDecoder(r.Body).Decode(&se)
		check(err)
		//full query using contians
		searchQuery := meetingplannerdb.QueryRow(`SELECT * FROM meetings WHERE userName=$1 AND agenda CONTAINS=$2`, userName, se)
		output(w, searchQuery)
	*/
}
