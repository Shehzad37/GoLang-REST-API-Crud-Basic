package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {

	users = append(users, User{"Tom", "Tom@cartoon.com"})
	users = append(users, User{"Jerry", "Jerry@cartoon.com"})
	users = append(users, User{"Loki", "loki@asgard.org"})
	users = append(users, User{"Thor", "Thor@asgard.org"})

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {

	handleRequest()

}
