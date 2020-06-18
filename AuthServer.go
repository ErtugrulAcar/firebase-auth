package main

import (
	"FirebaseAuth/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/user/create", handler.CreateUserHandler).Methods("POST")
	mux.HandleFunc("/user/signin", handler.SignIn).Methods("POST")
	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal(error)
	}

}
