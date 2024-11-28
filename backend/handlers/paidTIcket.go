package handlers

import (
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	event "local-event-management-backend/models/event"
	models "local-event-management-backend/models/payment"

	"local-event-management-backend/types"
	"net/http"
)


func writeError(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(types.GraphQLError{Message: message})
}


func PaidTicket(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}
	var payload types.TicketActionPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	event, err := event.GetEventById(payload.Input.Event_id)
	if err != nil{
		httpHelper.WriteError(w,http.StatusInternalServerError,"failed to get event")
		return
	}

	if len(event.Data_events) == 0  {
		httpHelper.WriteError(w,http.StatusBadRequest,"event not found")
		return		
	}

	if event.Data_events[0].Totalavailabletickets <= 0{
		httpHelper.WriteError(w,http.StatusBadRequest, "no avaliable tickets" )
		return
	}

	if !(event.Data_events[0].Ticketprice > 0) {
		httpHelper.WriteError(w,http.StatusBadRequest, "no free tickets" )
		return

	}

	response, err := models.InitiatePayment(event.Data_events[0].Ticketprice , payload.Input)
	if err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,  "unable to proccess the payment request" )
		return
	}
	json.NewEncoder(w).Encode(response)
	
}