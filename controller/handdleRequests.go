package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", getUsers).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8081", router))

}

