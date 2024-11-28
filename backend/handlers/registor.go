package handlers

import (
	"encoding/json"
	"io"
	"local-event-management-backend/helpers/auth"
	httpHelper "local-event-management-backend/helpers/http"
	models "local-event-management-backend/models/user"
	"local-event-management-backend/types"
	"net/http"
	"os"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	var payload types.InsertUserPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	
	result, err := models.InsertUser(payload.Input)
	if err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,"failed to create user")
		return
	}

    token, err := auth.MakeJWT(result.Insert_data_users_one.User_id, []byte(os.Getenv("JWT_KEY")), result.Insert_data_users_one.Role, time.Hour*24)
    if err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,"failed to login user")
		return
	}
	


    http.SetCookie(w, auth.SetCookie(token, 86400))

	res := types.LoginResponse{
		Success: true,
	}
	json.NewEncoder(w).Encode(res)
}
