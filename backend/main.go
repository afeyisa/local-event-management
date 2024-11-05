package main

import (
	"fmt"
	"local-event-management-backend/actions"
	"local-event-management-backend/auth"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)
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
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	mux.HandleFunc("/upload", actions.UploadImage)
	mux.HandleFunc("/tickectcheckout", actions.Ticketcheckout)
	fmt.Println("server started")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}