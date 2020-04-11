// main
// Simple API server

package main

import (
	"fmt"
	"mycircleapi/db"
	handler "mycircleapi/handlers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//initialize dynammo DB
	const region = "us-west-2"
	const endpoint = "" //"http://localhost:8080"
	db.InitializeDB(region, endpoint)

	router := mux.NewRouter().StrictSlash(true)
	setroutes(router)
	port := ":8080"
	log.Printf("API server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Setup server routes
// Setting server datapoints handllers routes
func setroutes(router *mux.Router) {

	usr := handler.NewUserController()
	userLogin := new(handler.UserController)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/login", userLogin.Login).Methods("POST")
	router.HandleFunc("/user/{id}", usr.Get).Methods("GET")
	router.HandleFunc("/user", usr.Set).Methods("POST")

	msg := handler.NewMsgController()
	//Messages
	router.HandleFunc("/message/{id}", msg.Get).Methods("GET")
	router.HandleFunc("/message", msg.Set).Methods("POST")

	//Contacts
	cont := handler.NewContactsController()
	router.HandleFunc("/contact", cont.Set).Methods("POST")
	router.HandleFunc("/contact/{id}", cont.Get).Methods("GET")
	router.HandleFunc("/contact/{id}", cont.Delete).Methods("DELETE")
	router.HandleFunc("/contact/{id}", cont.Update).Methods("PATCH")

	//registration
	reg := handler.NewRegistrationController()
	router.HandleFunc("/registration", reg.Set).Methods("POST")
	router.HandleFunc("/registration/{id}", reg.Get).Methods("GET")
	router.HandleFunc("/registration/{id}", reg.Delete).Methods("DELETE")
	router.HandleFunc("/registration/{id}", reg.Update).Methods("PATCH")

	//affected
	affect := handler.NewAffectedController()
	router.HandleFunc("/affectedlist", affect.Set).Methods("POST")
	router.HandleFunc("/affectedlist/{id}", affect.Get).Methods("GET")
	router.HandleFunc("/affectedlist/{id}", affect.Delete).Methods("DELETE")
	router.HandleFunc("/affectedlist/{id}", affect.Set).Methods("PATCH")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to MyCircle API service.\n Health: Ok")
}
