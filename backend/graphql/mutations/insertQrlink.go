package mutations

type Isert_qr_link struct{
	Update_data_tickets struct{
		Affected_rows int `json:"affected_rows"`
	}`graphql:"	update_data_tickets(where: {ticket_id: {_eq: $ticket_id}}, _set: {qr_link: $qr_link})"`
}

// mutation {
// 	update_data_tickets(where: {ticket_id: {_eq: ""}}, _set: {qr_link: ""}) {
// 	  affected_rows
// 	}
//   }