package models

import (
	"context"
	graphqlclient "local-event-management-backend/graphql/graphqlclient"
	queries "local-event-management-backend/graphql/queries"

	u "github.com/google/uuid"
)

func IsUserExsist(userId u.UUID )(bool, error){

	type uuid struct {
		u.UUID
	}

	variables := map[string]interface{}{
		"user_id":uuid{UUID: userId},
	}

	var res queries.UsersById
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return false , err
	}

	err = client.Query(context.Background(),&res,variables)
	if err != nil{
		return false , err
	}
	return len(res.Data_users) > 0, nil

}