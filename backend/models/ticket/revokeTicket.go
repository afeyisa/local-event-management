package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/mutations"

)

func RevokeTicket(variables map[string]interface{})(err error){
	
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return
	}
	var result mutations.DeleteTicket

	err = client.Mutate(context.Background(), &result, variables)
	
	if err != nil{
		fmt.Println("error revoking ticket", err)
		return
	}
	
	return nil
}
