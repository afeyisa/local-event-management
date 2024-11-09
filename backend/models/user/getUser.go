package models

import (
	"context"

	graphqlclient "local-event-management-backend/graphql/graphqlclient"
	queries "local-event-management-backend/graphql/queries"
)



func FetchHashedPasswordWithIdByEmail(email string) (queries.Users, error) {

	variables := map[string]interface{}{
		"email":  email,
	}
	
	var res queries.Users
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return res , err
	}

	err = client.Query(context.Background(),&res,variables)
	if err != nil{
		return res , err
	}

	return res, nil
}