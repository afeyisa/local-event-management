package types

import "github.com/google/uuid"

type TicketchekoutArgs struct {
	Event_id uuid.UUID `json:"event_id"`
	User_id  uuid.UUID `json:"user_id"`
    First_name string `json:"first_name,omitempty"`
    Last_name string `json:"last_name,omitempty"`
    Email string `json:"email,omitempty"`
    Phone_number string `json:"phone_number,omitempty"`
}

type Ticketcheckout struct {
	Url string `json:"url"`
}

type TicketActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            TicketchekoutArgs      `json:"input"`
}

type ChapaResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}
