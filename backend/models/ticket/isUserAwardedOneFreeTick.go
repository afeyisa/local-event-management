package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/queries"

	u "github.com/google/uuid"
)

func IsUserAwardedOneFreeTick(user_id, event_id u.UUID) (bool, error){
	type uuid struct {
		u.UUID
	}
	variables := map[string]interface{}{
		"ticket_owner_id": uuid{UUID: user_id},
		"event_id":uuid{UUID: event_id},
	}

	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return false, err
	}
	var result queries.TicketOwner

	err = client.Query(context.Background(), &result, variables)

	if err != nil {
		fmt.Println("error chacking ticket isawarded", err)
		return false, err
	}
	return len(result.Data_tickets)>0, nil
	
}