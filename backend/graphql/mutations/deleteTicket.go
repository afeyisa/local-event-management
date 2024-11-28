package mutations

type DeleteTicket struct{
	Delete_data_tickets struct{
		Affected_rows int `json:"affected_rows"`
	}`graphql:"delete_data_tickets(where: {tx_ref: {_eq: $tx_ref}})"`
}

