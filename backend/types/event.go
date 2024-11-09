
package types

type Events struct{
	Ticket_price float64 `json:"ticket_price"`
	Total_available_tickets int `json:"total_available_tickets"`
}

type EventBody struct {
	Data struct {
		Events []Events `json:"data_events"`
	} `json:"data"`
	Errors []GraphQLError `json:"errors"`
}