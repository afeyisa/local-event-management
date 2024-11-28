package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/queries"

	u "github.com/google/uuid"
)

func isTicketAwarded(tx_ref u.UUID) (bool, error){
	type uuid struct {
		u.UUID
	}
	variables := map[string]interface{}{
		"tx_ref": uuid{UUID: tx_ref},
	}

	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return false, err
	}
	var result queries.Ticket

	err = client.Query(context.Background(), &result, variables)

	if err != nil {
		fmt.Println("error chacking ticket isawarded", err)
		return false, err
	}
	return len(result.Data_tickets)>0, nil
	
}