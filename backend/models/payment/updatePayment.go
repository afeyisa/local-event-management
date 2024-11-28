package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/mutations"

	u "github.com/google/uuid"
)


func UPdatePaymentStatus(data Data) (result mutations.UdatePayment, err error){
	type uuid struct {
		u.UUID
	}
	tx, err:= u.Parse(data.TxRef)

	if err != nil{
		return
	}
	tx_refuuid := uuid{UUID:tx}

	// var paymentMethod *string
	variables := map[string]interface{}{
		"tx_ref":tx_refuuid,
		"chapa_reference": data.Reference,
		"chapa_charge": data.Charge,
		"status":data.Status,
		// "payment_method": paymentMethod,
	}
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return
	}
	// var result mutations.UdatePayment

	err = client.Mutate(context.Background(), &result, variables)
	
	if err != nil{
		fmt.Println("error ", err)
		return
	}


	return result , nil
}


func UPdatePaymentMethod(TxRef, paymentMethod string) (err error){

	type uuid struct {
		u.UUID
	}
	tx, err:= u.Parse(TxRef)

	if err != nil{
		return
	}
	tx_refuuid := uuid{UUID:tx}
	variables := map[string]interface{}{
		"tx_ref":tx_refuuid,
		"payment_method": paymentMethod,
	}
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil{
		return
	}

	var result mutations.UpdatePyamentMethod
	err = client.Mutate(context.Background(), &result, variables)
	
	if err != nil{
		fmt.Println("error udating payment method ", err)
		return
	}
	return nil
}