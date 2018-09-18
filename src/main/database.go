package main

import "database/sql"

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
		meetings (
			id SERIAL PRIMARY KEY,
			topic VARCHAR(20),
			dateAndTime TIMESTAMP, 
			roomID VARCHAR(5), 
			agenda VARCHAR(400),
			ownerID INT REFERENCES users (id)
			
			
		)`)
	check(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS 
		rooms (
			id SERIAL PRIMARY KEY,
			name VARCHAR(20)
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
