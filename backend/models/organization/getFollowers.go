package models

import (
	"context"

	graphqlclient "local-event-management-backend/graphql/graphqlclient"
	queries "local-event-management-backend/graphql/queries"

	u "github.com/google/uuid"
)

func GetFollower( orgId u.UUID ) (data queries.Data_organization_followers, err error){
	type uuid struct {
		u.UUID
	}

	variables := map[string]interface{}{
		"organization_id": uuid{UUID: orgId},
	}
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return
	}
	

	err = client.Query(context.Background(), &data, variables)

	if err != nil {
		return
	}
	return data , nil
}