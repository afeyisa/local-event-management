package main

import (

	"fmt"
	"log"
	"net/http"

	"local-event-management-backend/handlers"
	"github.com/joho/godotenv"
)


func main() {
	

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
		
	mux := http.NewServeMux()
 	mux.HandleFunc("/login", handlers.LoginHandler)
  	mux.HandleFunc("/register",handlers.RegisterHandler)
  	mux.HandleFunc("/auth",handlers.HasuraAuthHandler)
	mux.HandleFunc("/logout", handlers.LogOutHandler)
	mux.HandleFunc("/protectpageroute",handlers.ProtectPageRoute)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	mux.HandleFunc("/upload", handlers.UploadImage)
	mux.HandleFunc("/tickectcheckout", handlers.Ticketcheckout)
	fmt.Println("server started")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}