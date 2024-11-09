package handlers

import (
	"encoding/json"
	"io"
	models "local-event-management-backend/models/ticket"
	"local-event-management-backend/types"
	"net/http"
)

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
	var payload types.TicketActionPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Send the request params to the Action's generated handler function
	result, err := models.Ticketchekout(payload.Input)

	// throw if an error happens
	if err != nil {
		errorObject := types.GraphQLError{
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