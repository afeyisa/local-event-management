package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/mutations"
	"local-event-management-backend/types"

	u "github.com/google/uuid"
)

func InsertPayments(price_amount float64,  userInfo types.TicketchekoutArgs ,tx_ref string) (res string, err error) {
	type uuid struct {
		u.UUID
	}
	
	eventUUID := uuid{UUID: userInfo.Event_id}
	userUUID := uuid{UUID: userInfo.User_id}
	tx_refUUID, err := u.Parse(tx_ref)

	if err != nil{
		return
	}

	tx_refuuid := uuid{UUID:tx_refUUID }
	variables := map[string]interface{}{
		"amount": price_amount,
		"currency": "ETB", 
		"email": userInfo.Email, 
		"first_name": userInfo.First_name, 
		"last_name": userInfo.Last_name,
		"paid_for_event_id": eventUUID, 
		"payer_user_id": userUUID, 
		"phone_number": userInfo.Phone_number, 
		"tx_ref": tx_refuuid,
		"status": "initialized",
		}
	 fmt.Println(variables["phone_number"])
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	var result mutations.InsetPayment
	err = client.Mutate(context.Background(), &result, variables)
	if err != nil {
		return
	}
	return result.Insert_data_payments_one.Tx_ref.String(), nil
}