package email

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)



func GetGmailClient(tokenFilePath string, config *oauth2.Config) (client *http.Client, err error) {

	token, err := tokenFromFile(tokenFilePath)
	if err != nil {
		return
	}
	tokenSource := config.TokenSource(context.Background(), token)
	reUseTokenSource := oauth2.ReuseTokenSource(token, tokenSource)
	return oauth2.NewClient(context.Background(), reUseTokenSource), nil
}
