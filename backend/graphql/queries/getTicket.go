package queries

import "github.com/google/uuid"

type Ticket struct {
	Data_tickets []struct {
		Ticket_owner_id uuid.UUID `json:"ticket_owner_id"`
		Tx_ref uuid.UUID `json:"tx_ref"`
		Event_id uuid.UUID `json:"event_id"`
	}`graphql:"data_tickets(where: {tx_ref: {_eq: $tx_ref}})"`
}
