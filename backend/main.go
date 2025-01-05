package main

import (
	"fmt"
	"local-event-management-backend/helpers/auth"
	"local-event-management-backend/handlers"
	"local-event-management-backend/middleware"
	email "local-event-management-backend/models/email"
	// ticket "local-event-management-backend/models/ticket"
	"log"
	"net/http"
	"os"
	// "syscall"

	// "github.com/joho/godotenv"
	// "golang.org/x/term"
	"google.golang.org/api/gmail/v1"
)




func main() {
		
	// err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }
	
	// fmt.Print("Please Credential key: ")
	// k, err := term.ReadPassword(int(syscall.Stdin))
	// if err != nil {
	// 	log.Fatalf("Failed to read key: %v", err)
	// }

	// fmt.Println()
	f := "credentials"

	email.InitializeGmailService(auth.DecryptCredentials(f, os.Getenv("CRIDENTIALS")),"token.json",gmail.GmailSendScope)
	mux := http.NewServeMux()
 	mux.Handle("/login",					 middleware.HasuraAuth(http.HandlerFunc(handlers.LoginHandler)))
  	mux.Handle("/register",					 middleware.HasuraAuth(http.HandlerFunc(handlers.RegisterHandler)))
	mux.Handle("/logout", 				     middleware.HasuraAuth(http.HandlerFunc(handlers.LogOutHandler)))
	mux.Handle("/protectpageroute",			 middleware.HasuraAuth(http.HandlerFunc(handlers.ProtectPageRoute)))
	mux.Handle("/upload", 					 middleware.HasuraAuth(http.HandlerFunc(handlers.ImageUploads)))
	mux.Handle("/tickectcheckout/paid", 	 middleware.HasuraAuth(middleware.TicketCheckOutAuth(http.HandlerFunc(handlers.PaidTicket))))
	mux.Handle("/tickectcheckout/free", 	 middleware.HasuraAuth(middleware.TicketCheckOutAuth(middleware.FreeTicketAuth(http.HandlerFunc(handlers.FreeTicket)))))
	mux.Handle("/serveimage", 				 middleware.HasuraAuth(http.HandlerFunc(handlers.ServeImage)))
	mux.Handle("/deletethumbnailimage",		 middleware.HasuraAuth(http.HandlerFunc(handlers.DeleteThumbnail)))
	mux.Handle("/deletetotherimages",		 middleware.HasuraAuth(http.HandlerFunc(handlers.DeleteImages)))
	mux.Handle("/verifypayment", 			 middleware.HasuraAuth(http.HandlerFunc(handlers.VerifyPaymentHandler)))
	mux.Handle("/emailsender",               middleware.HasuraAuth(http.HandlerFunc(handlers.SendEmalToFollowers)))
	
	mux.Handle("/chapawebhook", 			 middleware.ChapaAuth(http.HandlerFunc(handlers.ChapaWebHookHandler)))

	fmt.Println("server started:")
	err = http.ListenAndServe(os.Getenv("PORT"), mux)
	log.Fatal(err)
}