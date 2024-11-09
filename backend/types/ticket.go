package types



type TicketchekoutArgs struct{
	Event_id string `json:"event_id"`
	User_id string `json:"user_id"`
}

type Ticketcheckout struct {
	Url string `json:"url"`
  }

type TicketActionPayload struct {
  SessionVariables map[string]interface{} `json:"session_variables"`
  Input            TicketchekoutArgs `json:"input"`
}


type ChapaResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data struct {
	  CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
  }
