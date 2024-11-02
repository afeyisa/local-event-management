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
	// 	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create upload directory:", err)
	}

	// Serve the uploads directory
	// http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadPath))))

	// Handle the upload request
	
	
	mux := http.NewServeMux()
 	mux.HandleFunc("/login", actions.LoginHandler)
  	mux.HandleFunc("/register",actions.RegisterHandler)
  	mux.HandleFunc("/auth",auth.HasuraAuthHandler)
	mux.HandleFunc("/logout", actions.LogOutHandler)
	mux.HandleFunc("/protectpageroute",actions.ProtectPageRoute)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	mux.HandleFunc("/upload", actions.UploadImage)
	mux.HandleFunc("/tickectcheckout", actions.Ticketcheckout)

	// mux.HandleFunc("/uploads",actions.ServeImage)
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}