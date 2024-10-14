package main

import (
	"local-event-management-backend/actions"
	"local-event-management-backend/auth"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)
// HTTP server for the handler
func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	mux := http.NewServeMux()
 	mux.HandleFunc("/login", actions.LoginHandler)
  	mux.HandleFunc("/register",actions.ResgisterHandler)
  	mux.HandleFunc("/auth",auth.AuthHandler)
	mux.HandleFunc("/logout", actions.LogOutHandler)

	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}