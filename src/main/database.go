package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

func testOverviews(meetingplannerdb *sql.DB) {
	users, err := meetingplannerdb.Query(`SELECT * FROM users`)
	check(err)

	defer users.Close()

	fmt.Println("Users: \n ")

	for users.Next() {
		var (
			id       int
			name     string
			phone    string
			email    string
			password string
		)

		err := users.Scan(&id, &name, &phone, &email, &password)
		check(err)

		fmt.Println(id, name, phone, email, password)
	}

	rooms, err := meetingplannerdb.Query(`SELECT * FROM rooms`)
	check(err)

	defer rooms.Close()

	fmt.Println("Rooms: \n ")

	for rooms.Next() {
		var (
			id   int
			name string
		)

		err := rooms.Scan(&id, &name)
		check(err)

		fmt.Println(id, name)
	}

	meetings, err := meetingplannerdb.Query(`SELECT * FROM meetings`)
	check(err)

	defer meetings.Close()

	fmt.Println("Meetings: \n ")

	for meetings.Next() {
		var (
			id          int
			topic       string
			dateAndTime time.Time
			agenda      string
			roomID      int
			ownerID     int
		)

		err := meetings.Scan(&id, &topic, &dateAndTime, &agenda, &roomID, &ownerID)
		check(err)

		fmt.Println(id, topic, dateAndTime, agenda, roomID, ownerID)
	}

	priorMeetings, err := meetingplannerdb.Query(`SELECT * FROM priorMeetings`)
	check(err)

	defer priorMeetings.Close()

	fmt.Println("PriorMeetings: \n ")

	for priorMeetings.Next() {
		var (
			id        int
			userID    int
			meetingID int
		)

		err := priorMeetings.Scan(&id, &meetingID, &userID)
		check(err)

		fmt.Println(id, meetingID, userID)
	}

	participants, err := meetingplannerdb.Query(`SELECT * FROM participants`)
	check(err)

	defer participants.Close()

	fmt.Println("Participants: \n ")

	for participants.Next() {
		var (
			id        int
			userID    int
			meetingID int
		)

		err := participants.Scan(&id, &meetingID, &userID)
		check(err)

		fmt.Println(id, meetingID, userID)
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

	_, err = tx.Exec("DROP SCHEMA IF EXISTS public CASCADE")
	_, err = tx.Exec("CREATE SCHEMA public")
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50), 
			phone VARCHAR(20), 
			email VARCHAR(50),
			password VARCHAR(20)
		)`)
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		rooms (
			id SERIAL PRIMARY KEY,
			name VARCHAR(20)
		)`)
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		meetings (
			id SERIAL PRIMARY KEY,
			topic VARCHAR(20),
			dateAndTime TIMESTAMP, 
			agenda VARCHAR(1000),
			roomID INT REFERENCES rooms(id), 
			ownerID INT REFERENCES users (id)
		)`)
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		priorMeetings (
			id SERIAL PRIMARY KEY,
			meetingID INT REFERENCES users (id),
			userID INT REFERENCES users (id)
		)`)
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		participants (
			id SERIAL PRIMARY KEY,
			meetingID INT REFERENCES users (id),
			userID INT REFERENCES users (id)
		)`)
	check(err)

	tx.Commit()
}

//
// func init() {
//     rand.Seed(time.Now().UnixNano())
// }

// Random letter seeder
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Seed database with arbitary data
func seed(db *sql.DB) {
	tx, err := db.Begin()
	check(err)

	for i := 1; i < 11; i++ {

		// Users seed
		_, err = tx.Exec(`INSERT INTO 
		users(name, phone, email, password) VALUES(
			$1, $2, $3, $4)`, randStr(15), randStr(15), randStr(15), randStr(15))
		check(err)

		// Rooms seed
		_, err = tx.Exec(`INSERT INTO
		rooms(name) VALUES(
			$1)`, randStr(15))
		check(err)

		// Meetings seed
		_, err = tx.Exec(`INSERT INTO
		meetings(topic, dateAndTime, agenda, roomID, ownerID) VALUES(
			$1, $2, $3, $4, $5)`, randStr(15), "2001-09-28 01:00", randStr(15), i, 1)
		check(err)

		// PriorMeetings seed
		_, err = tx.Exec(`INSERT INTO
		priorMeetings(meetingID, userID) VALUES(
			$1, $2)`, i, i)
		check(err)

		// Participants seed
		_, err = tx.Exec(`INSERT INTO
		participants(meetingID, userID) VALUES(
			$1, $2)`, i, i)
		check(err)
	}

	tx.Commit()
}
