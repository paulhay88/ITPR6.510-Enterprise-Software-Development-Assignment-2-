package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

// Check for errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openDatabase() *sql.DB {

	db, err := sql.Open("postgres", "user=postgres password=password dbname=meetingplannerdb sslmode=disable")
	check(err)

	err = db.Ping()
	check(err)

	return db
}

func createTables(db *sql.DB) {


	tx, err := db.Begin()
	check(err)

	_, err = tx.Exec("DROP TABLE IF EXISTS users")
	_, err = tx.Exec("DROP TABLE IF EXISTS rooms")
	_, err = tx.Exec("DROP TABLE IF EXISTS meetings")
	_, err = tx.Exec("DROP TABLE IF EXISTS meetingDetails")
	_, err = tx.Exec("DROP TABLE IF EXISTS participants")

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		users(
			id SERIAL INT PRIMARY KEY, 
			name VARCHAR(50), 
			phone VARCHAR(20), 
			email VARCHAR(50),
			password VARCHAR(20)
		)`)

	// exec(tx, "CREATE TABLE IF NOT EXISTS companies(index INT PRIMARY KEY, company VARCHAR(50), staff_size INT, address VARCHAR(100), phone VARCHAR(20), email VARCHAR(50))")
	// exec(tx, "CREATE TABLE IF NOT EXISTS interns(index SERIAL PRIMARY KEY, first_name VARCHAR(25), last_name VARCHAR(25), phone VARCHAR(20), email VARCHAR(50), company_id INT)")
	// SERIAL use is postgres specific. Would usually be auto_increment

	tx.Commit()
}

func main() {
	// defer profile.Start().Stop()

	meetingsDB := openDatabase()

	// Defer closing of database to end of main()
	defer meetingsDB.Close()


	createTables(meetingsDB)

	overview, err := meetingsDB.Query(`SELECT * FROM users`)
	check(err)

	defer overview.Close()

	for overview.Next() {
		var (
			name    string
			phone string
			email string
			password string
		)

		err := overview.Scan(&name, &phone, &email, &password)
		check(err)

		fmt.Println(name, phone, email, password)
	}


	fmt.Println(meetingsDB)
}
