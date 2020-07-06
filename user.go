package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `josn:"email"`
}

var users []User

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from Homepage")

}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "All Users")

	json.NewEncoder(w).Encode(users)

}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "New User")

	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	user.Name = "Vero"
	user.Email = "Vero@mail.com"
	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]

	for index, item := range users {
		if item.Email == email {
			users = append(users[:index], users[index+1:]...)
			var user User

			_ = json.NewDecoder(r.Body).Decode(&user)

			user.Name = item.Name
			user.Email = "NewMail.com"
			users = append(users, user)
			//json.NewEncoder(w).Encode(user)
			break
		}

	}

	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := mux.Vars(r)["name"]

	for index, item := range users {

		if item.Name == name {

			users = append(users[:index], users[index+1:]...)
			break

		}

	}

	json.NewEncoder(w).Encode(users)

}
