package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
type Events struct{
	Ticket_price float64 `json:"ticket_price"`
	Total_available_tickets int `json:"total_available_tickets"`
}

type EventBody struct {
	Data struct {
		Events []Events `json:"data_events"`
	} `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

type ticketchekoutArgs struct{
	Event_id string `json:"event_id"`
	User_id string `json:"user_id"`
}

type ticketcheckout struct {
	Url string `json:"url"`
  }
type TicketActionPayload struct {
  SessionVariables map[string]interface{} `json:"session_variables"`
  Input            ticketchekoutArgs `json:"input"`
}


type ChapaResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data struct {
	  CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
  }

func Ticketcheckout(w http.ResponseWriter, r *http.Request) {

  // set the response header as JSON
  w.Header().Set("Content-Type", "application/json")

  // read request body
  reqBody, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  // parse the body as action payload
  var actionPayload TicketActionPayload
  err = json.Unmarshal(reqBody, &actionPayload)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  // Send the request params to the Action's generated handler function
  result, err := ticketchekout(actionPayload.Input)

  // throw if an error happens
  if err != nil {
    errorObject := GraphQLError{
      Message: err.Error(),
    }
    errorBody, _ := json.Marshal(errorObject)
    w.WriteHeader(http.StatusBadRequest)
    w.Write(errorBody)
    return
  }

  // Write the response as JSON
  data, _ := json.Marshal(result)
  w.Write(data)

}

// Auto-generated function that takes the Action parameters and must return it's response type
func ticketchekout(args ticketchekoutArgs) (response ticketcheckout, err error) {
	query := fmt.Sprintf(`{
		data_events(where: {event_id: {_eq: "%s"}}) {
			ticket_price
			total_available_tickets
		}
	}`, args.Event_id)

	fmt.Println(query)
	reqBody := GraphQLRequest{
		Query: query,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/graphql", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return ticketcheckout{}, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return ticketcheckout{}, fmt.Errorf("failed to request event data")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ticketcheckout{}, fmt.Errorf("failed to fetch event data")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ticketcheckout{}, fmt.Errorf("failed to read event data")
	}

	event := EventBody{}
	fmt.Println(event)

	if err := json.Unmarshal(respBody, &event); err != nil {
		return ticketcheckout{}, fmt.Errorf("failed to get event")
	}

	if len(event.Errors) > 0 {
		return ticketcheckout{},  fmt.Errorf("failed to get event")
	}

	

	if len(event.Data.Events) == 0 {

		return ticketcheckout{},  fmt.Errorf("failed to get event")
	}

	if (event.Data.Events[0].Ticket_price>0 && event.Data.Events[0].Total_available_tickets>0){
	
	url := "https://api.chapa.co/v1/transaction/initialize"

	payload := map[string]interface{}{
		"amount":        event.Data.Events[0].Ticket_price,
		"currency":      "ETB",
		"user_id":       args.User_id,
	  }

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return ticketcheckout{}, err
	}
	  req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	  if err != nil {
		return ticketcheckout{}, err
	  }
	  req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("CHAPA")))
	  req.Header.Set("Content-Type", "application/json")
		  
	  res, err := client.Do(req)
	  if err != nil {
		return ticketcheckout{}, err
	  }
	  defer res.Body.Close()
	
	  var chapaResponse ChapaResponse
	  if err := json.NewDecoder(res.Body).Decode(&chapaResponse); err != nil {
		return ticketcheckout{}, err
	  }
	
	  if chapaResponse.Status != "success" {
		return ticketcheckout{}, fmt.Errorf("chapa API error: %s", chapaResponse.Message)
	  }
	
  response =  ticketcheckout {
    Url: chapaResponse.Data.CheckoutURL,
  }
  return response, nil
}
return ticketcheckout{}, fmt.Errorf("no avaliable tickets or free event")
}
