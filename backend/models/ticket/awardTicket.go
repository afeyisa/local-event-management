package models

import (
	"context"
	"fmt"
	"local-event-management-backend/graphql/graphqlclient"
	"local-event-management-backend/graphql/mutations"
	"os"

	u "github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

func Insert_ticket_qr(link string, ticket_id u.UUID) error{
	type uuid struct {
		u.UUID
	}
	variables := map[string]interface{}{
		"ticket_id": uuid{UUID: ticket_id},
		"qr_link":link,
	}

	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return  err
	}
	var result mutations.Isert_qr_link

	err = client.Mutate(context.Background(), &result, variables)

	if err != nil {
		fmt.Println("error inserting ticket qr", err)
		return  err
	}
	fmt.Println(result.Update_data_tickets.Affected_rows)
	return nil
	
}

func AwardPaidTicket(ticket map[string]interface{}) (err error) {
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return
	}
	var result mutations.CreateTPaidicket

	err = client.Mutate(context.Background(), &result, ticket)

	if err != nil {
		fmt.Println("error awarding ticket", err)
		return
	}

	// generate qr code for ticket 
	qr_name := result.Insert_data_tickets_one.Ticket_id.String()+".png"
	err = qrcode.WriteFile(os.Getenv("TICKET_VERIFICATION_LINK")+result.Insert_data_tickets_one.Ticket_id.String(), qrcode.High, 256, os.Getenv("UPLOAD_PATH")+qr_name)
	if err != nil{
		return
	}
	err = Insert_ticket_qr(os.Getenv("BASE_URL")+"/"+qr_name, result.Insert_data_tickets_one.Ticket_id)
	if err != nil {
		return
	}

	return nil
}

func AwardFreeTicket(ticket map[string]interface{}) (err error) {
	client, err := graphqlclient.NewGraphqlClient()
	if err != nil {
		return
	}
	var result mutations.CreateFreeTicket

	err = client.Mutate(context.Background(), &result, ticket)

	if err != nil {
		fmt.Println("error awarding ticket", err)
		return
	}

	// generate qr code for ticket 
	qr_name := result.Insert_data_tickets_one.Ticket_id.String()+".png"
	err = qrcode.WriteFile(os.Getenv("TICKET_VERIFICATION_LINK")+result.Insert_data_tickets_one.Ticket_id.String(), qrcode.High, 256, os.Getenv("UPLOAD_PATH")+qr_name)
	if err != nil{
		return
	}
	err = Insert_ticket_qr(os.Getenv("BASE_URL")+"/"+qr_name, result.Insert_data_tickets_one.Ticket_id)
	if err != nil {
		return
	}

	return nil
}