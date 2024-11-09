package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/auth"
	models "local-event-management-backend/models/user"
	"local-event-management-backend/types"
	"net/http"
	"os"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// parse the body as action payload
	var payload types.InsertUserPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	
	result, err := models.InsertUser(payload.Input)
	if err != nil {
		errorObject := types.GraphQLError{
			Message: "un able to create user with " + payload.Input.Email,
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

    token, err := auth.MakeJWT(result.Insert_data_users_one.User_id, []byte(os.Getenv("JWT_KEY")), result.Insert_data_users_one.Role, time.Hour*24)
    if err != nil {
		errorObject := types.GraphQLError{
			Message: "unable login",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}
	


    http.SetCookie(w, auth.SetCookie(token, 86400))

	res := types.LoginResponse{
		Success: true,
	}
	data, _ := json.Marshal(res)
	w.Write(data)
}
