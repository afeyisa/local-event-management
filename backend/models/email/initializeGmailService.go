package email

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var GmailSeverce *gmail.Service

func InitializeGmailService(credentials []byte, token, scope string){
	ctx := context.Background()
	config, err := google.ConfigFromJSON(credentials, scope)
	
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, err := GetGmailClient(token,config)

	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}
	GmailSeverce, err = gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail service: %v", err)
	}
	
}