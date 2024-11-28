package handlers

import (
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	event "local-event-management-backend/models/event"
	ticket "local-event-management-backend/models/ticket"
	"local-event-management-backend/types"
	"net/http"

	u "github.com/google/uuid"
)

func FreeTicket(w http.ResponseWriter, r *http.Request) {
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

    //
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

    if event.Data_events[0].Ticketprice > 0{
		httpHelper.WriteError(w,http.StatusBadRequest, "This is not a free event" )
		return
	}


	type uuid struct {
		u.UUID
	}
	
	err = ticket.AwardFreeTicket(map[string]interface{}{
		"event_id":uuid{UUID:payload.Input.Event_id },
		"ticket_owner_id": uuid{UUID: payload.Input.User_id },
	})

	if err != nil {
		writeError(w,http.StatusInternalServerError,  "unable to proccess the ticket request" )
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"you got free ticket"})
}