package handlers

import (
	"encoding/json"
	"local-event-management-backend/auth"
	"local-event-management-backend/types"
	"net/http"
)



func LogOutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, auth.SetCookie("",-86400))
	res := types.LogOutResponse{Success: true}
	data, _ := json.Marshal(res)
    w.Write(data)
}