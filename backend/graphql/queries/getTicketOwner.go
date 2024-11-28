package queries

import "github.com/google/uuid"
type TicketOwner struct{
	Data_tickets []struct{
		Ticket_id uuid.UUID  `json:"event_id"`
		Ticket_owner_id  uuid.UUID `json:"ticket_owner_id"`
	}`graphql:"data_tickets(where: {_and: {event_id: {_eq: $event_id}, ticket_owner_id: {_eq: $ticket_owner_id}}})"`

}

// {
// 	data_tickets(where: {ticket_owner_id: {_eq: ""}}) {
// 		ticket_owner_id
// 	}
//   }
  
// {
// 	data_tickets(where: {ticket_owner_id: {_eq: ""}}) {
// 		ticket_owner_id
// 	}
//   }
// (where: {_and: {event_id: {_eq: ""}, ticket_owner_id: {_eq: ""}}})

// Ticket_owner_id  uuid.UUID `json:"ticket_owner_id"`
// }`graphql:"data_tickets(where: {_and: {event_id: {_eq: ""}, ticket_owner_id: {_eq: ""}}})"`
