package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/queries"

	u "github.com/google/uuid"
)

func GetPaymentStatus(tx_ref string) (payment queries.DataPayment, err error){
	type uuid struct{
		u.UUID
	}
	tx, err := u.Parse(tx_ref)

	tx_uuid := uuid{UUID: tx}
	if err != nil{
		return
	}

	variables := map[string]interface{}{
		"tx_ref":tx_uuid,
	}
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return
	}
	var result queries.PaymentStatus
	err = client.Query(context.Background(), &result, variables)
	if err != nil{
		return
	}
	fmt.Println(len(result.Data_Payment))
	return result.Data_Payment, nil
}