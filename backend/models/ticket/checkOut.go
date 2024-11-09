package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/types"
	"net/http"
	"os"
)

func Ticketchekout(args types.TicketchekoutArgs) (response types.Ticketcheckout, err error) {
	query := fmt.Sprintf(`{
		data_events(where: {event_id: {_eq: "%s"}}) {
			ticket_price
			total_available_tickets
		}
	}`, args.Event_id)

	fmt.Println(query)
	reqBody := types.GraphQLRequest{
		Query: query,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/graphql", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return response, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return response, fmt.Errorf("failed to request event data")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("failed to fetch event data")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("failed to read event data")
	}

	event := types.EventBody{}
	fmt.Println(event)

	if err := json.Unmarshal(respBody, &event); err != nil {
		return response, fmt.Errorf("failed to get event")
	}

	if len(event.Errors) > 0 {
		return response, fmt.Errorf("failed to get event")
	}

	if len(event.Data.Events) == 0 {

		return response, fmt.Errorf("failed to get event")
	}

	if event.Data.Events[0].Ticket_price > 0 && event.Data.Events[0].Total_available_tickets > 0 {

		url := "https://api.chapa.co/v1/transaction/initialize"

		payload := map[string]interface{}{
			"amount":   event.Data.Events[0].Ticket_price,
			"currency": "ETB",
			"user_id":  args.User_id,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			return response, err
		}
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return response, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("CHAPA")))
		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			return response, err
		}
		defer res.Body.Close()

		var chapaResponse types.ChapaResponse
		if err := json.NewDecoder(res.Body).Decode(&chapaResponse); err != nil {
			return response, err
		}

		if chapaResponse.Status != "success" {
			return response, fmt.Errorf("chapa API error: %s", chapaResponse.Message)
		}

		response = types.Ticketcheckout{
			Url: chapaResponse.Data.CheckoutURL,
		}
		return response, nil
	}
	return response, fmt.Errorf("no avaliable tickets or free event")
}
