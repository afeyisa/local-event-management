package mutations

import "github.com/google/uuid"

type CreateTPaidicket struct {
	Insert_data_tickets_one struct{
	Ticket_id uuid.UUID `json:"ticket_id"`
}`graphql:"insert_data_tickets_one(object: {event_id: $event_id, ticket_owner_id: $ticket_owner_id, tx_ref: $tx_ref})"`
}

type CreateFreeTicket struct {
	Insert_data_tickets_one struct{
	Ticket_id uuid.UUID `json:"ticket_id"`
}`graphql:"insert_data_tickets_one(object: {event_id: $event_id, ticket_owner_id: $ticket_owner_id})"`
}

