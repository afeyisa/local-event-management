package models

import (

	"context"
	"fmt"

	"local-event-management-backend/auth"
	graphqlclient "local-event-management-backend/graphql/graphqlclient"
	mutations "local-event-management-backend/graphql/mutations"
	"local-event-management-backend/types"
)

func InsertUser(args types.InsertUserArgs) (res mutations.InsertUser, err error) {
	hash, err := auth.HashPassword(args.Password)
	if err != nil{
		return
	}
    variables := map[string]interface{}{
        "email":args.Email,
        "password": hash,
        "role":"user",

    }

    client,err := graphqlclient.NewGraphqlClient()
    if err != nil{
		return 
	}

	err = client.Mutate(context.Background(),&res,variables)
	if err != nil{
        fmt.Println(err)
		return 
	}

	return
}
