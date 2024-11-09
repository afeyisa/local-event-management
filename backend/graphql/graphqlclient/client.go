package graphqlclient

import (
	"fmt"
	"net/http"
	"os"

	`github.com/hasura/go-graphql-client`
)

type Doer struct {
	Client *http.Client
}

func (d *Doer) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))
	return d.Client.Do(req)
}


func NewGraphqlClient()(*graphql.Client, error){
	hasuraURL := os.Getenv("HASURA")
	if hasuraURL == "" {
		return nil, fmt.Errorf("HASURA environment variable is not found")
	}
	return graphql.NewClient(hasuraURL, &Doer{Client: &http.Client{}}), nil
	
}
