package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	queries "local-event-management-backend/graphql/queries"
	"local-event-management-backend/types"

	u "github.com/google/uuid"
)


func Ticketchekout(args types.TicketchekoutArgs) (response types.Ticketcheckout, err error) {

	type uuid struct {
		u.UUID
	}
	
	eventUUID := uuid{UUID: args.Event_id}

	variables := map[string]interface{}{
		"event_id": eventUUID,
	}

	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return response, fmt.Errorf("failed to query event data")
	}

	var event queries.Events
	err = client.Query(context.Background(), &event, variables)

	if err != nil {
		fmt.Println(err)
		return response, fmt.Errorf("failed to query event data")
	}

	if len(event.Data_events) == 0 {

		return response, fmt.Errorf("the event not found")
	}

	if event.Data_events[0].Ticketprice > 0 && event.Data_events[0].Totalavailabletickets > 0 {
		response, err = InitiatePayment(event.Data_events[0].Ticketprice , args)
		return

	}
	// do something
	return response, fmt.Errorf("no avaliable tickets or the event is free")
}


