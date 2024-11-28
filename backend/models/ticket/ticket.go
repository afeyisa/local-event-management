package models

import (
	"fmt"
	"local-event-management-backend/graphql/mutations"

	u "github.com/google/uuid"
)

// in this function  what we will do is
// 1 check if the ticket is checked out for the user
// 2 check the status of the payment
// check the status of the payment
//


func Ticket(p mutations.UdatePayment) (err error){
	type uuid struct {
		u.UUID
	}
	isawarded , err := isTicketAwarded(p.Update_data_payments_by_pk.Tx_ref)
	fmt.Println(isawarded)

	if err != nil{
		return
	}
	

	if p.Update_data_payments_by_pk.Status == "success"&& !isawarded{
		// award ticket
		variables := map[string]interface{}{
			"event_id":uuid{UUID: p.Update_data_payments_by_pk.Paid_for_event_id},
			"ticket_owner_id": uuid{UUID: p.Update_data_payments_by_pk.Payer_user_id},
			"tx_ref": uuid{UUID: p.Update_data_payments_by_pk.Tx_ref},
		}

		err = AwardPaidTicket(variables)

		if err != nil{
			return
		}
	}else if p.Update_data_payments_by_pk.Status != "success"&& isawarded{

		err = RevokeTicket(map[string]interface{}{"tx_ref": uuid{UUID: p.Update_data_payments_by_pk.Tx_ref},})
		if err != nil{
			return
		}

	}
	return nil
}