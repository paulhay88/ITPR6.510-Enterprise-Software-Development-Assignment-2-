package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Needs to be admin to edit rooms
func getRooms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var rooms Rooms

	// Get rooms
	results, err := meetingplannerdb.Query(`SELECT * FROM rooms`)
	check(err)
	defer results.Close()

	for results.Next() {
		var room Room

		err = results.Scan(&room.ID, &room.Name)
		check(err)

		rooms.Rooms = append(rooms.Rooms, room)
	}

	output(w, rooms.Rooms)
}

func createRoom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var room = new(Room)

	// User creating room
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Extract meeting from json
	err = json.NewDecoder(r.Body).Decode(&room)
	check(err)

	// Make sure only admin user can create
	if userName == "admin" {
		// Create meeting
		_, err = meetingplannerdb.Exec(
			`INSERT INTO rooms(name) VALUES($1)`,
			room.Name)
		check(err)
	}

}

func updateRoom(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var room Room

	// User updating room
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Decode user passed data
	err = json.NewDecoder(r.Body).Decode(&room)
	check(err)

	// Get room id
	room.ID, err = strconv.Atoi(params.ByName("id"))
	check(err)

	// Only admin can add rooms
	if userName == "admin" {
		// Update if name isn't null
		_, err = meetingplannerdb.Exec(
			`UPDATE rooms
		SET name = COALESCE(NULLIF($1, ''), name)
		WHERE
			id = $2`,
			room.Name, room.ID)
	}
}

func deleteRoom(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Get room id
	roomID, err := strconv.Atoi(params.ByName("id"))
	check(err)

	// User deleting room
	meetingCookie, err := r.Cookie("authUser")
	check(err)
	userName := strings.Split(meetingCookie.Value, ":")[0]

	// Only admin can delete rooms
	if userName == "admin" {
		_, err = meetingplannerdb.Exec(
			`DELETE FROM rooms
				WHERE id = $1`,
			roomID)
		check(err)
	}
}
