package main

import (
	"fmt"
	"local-event-management-backend/handlers"
	"local-event-management-backend/middleware"
	"log"
	"net/http"
	"os"

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
	mux.HandleFunc("/upload", handlers.ImageUploads)
	mux.HandleFunc("/tickectcheckout", handlers.Ticketcheckout)
	mux.HandleFunc("/serveimage", handlers.ServeImage)

	fmt.Println("server started")
	err = http.ListenAndServe(os.Getenv("PORT"), middleware.Auth(mux) )
	log.Fatal(err)
}