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
  	mux.HandleFunc("/register",actions.RegisterHandler)
  	mux.HandleFunc("/auth",auth.HasuraAuthHandler)
	mux.HandleFunc("/logout", actions.LogOutHandler)
	mux.HandleFunc("/protectpageroute",actions.ProtectPageRoute)
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}