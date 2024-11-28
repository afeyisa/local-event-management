package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	models "local-event-management-backend/models/ticket"
	"local-event-management-backend/types"
	"net/http"
)

// check if user has already awarded free ticket


func FreeTicketAuth(next http.Handler) http.Handler{

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

	isUserAwardedOneFreeTick, err := models.IsUserAwardedOneFreeTick(payload.Input.User_id, payload.Input.Event_id)

	if err != nil{
		httpHelper.WriteError(w,http.StatusInternalServerError,"error veryfing user")
		return
	}

	if isUserAwardedOneFreeTick{
		httpHelper.WriteError(w,http.StatusUnauthorized,"reached max ticket limit")
		return
	}
		
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	next.ServeHTTP(w,r)
},
)
}