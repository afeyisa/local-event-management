package queries

type Events struct {
	Data_events []struct {
		Ticketprice           float64 `graphql:"ticket_price"`
		Totalavailabletickets int     `graphql:"total_available_tickets"`
	} `graphql:"data_events(where: {event_id: {_eq: $event_id}})"`
}
