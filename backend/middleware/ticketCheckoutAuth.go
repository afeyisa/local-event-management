package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	httpHelper "local-event-management-backend/helpers/http"
	user "local-event-management-backend/models/user"
	"local-event-management-backend/types"
	"net/http"
)

func TicketCheckOutAuth(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	reqBody, err := io.ReadAll(r.Body)
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

	// NOTE
	// I SHOULD ADD USER VERIFICATION TO BUY TICKET 
	isexsist, err := user.IsUserExsist(payload.Input.User_id)

	if err != nil{
		httpHelper.WriteError(w,http.StatusInternalServerError,"error verifying user")
		return
	}

	if !isexsist{
		httpHelper.WriteError(w,http.StatusInternalServerError,"unauthorized")
		return
	}
		
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	next.ServeHTTP(w,r)
},
)
}
