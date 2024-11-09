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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing login")

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// parse the body as action payload
	var payload types.LoginPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		errorObject := types.GraphQLError{
			Message: "invalid payload",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return

	}

	resp, err := models.FetchHashedPasswordWithIdByEmail(payload.Input.Email)

	if err != nil {
		errorObject := types.GraphQLError{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	if len(resp.Data_users) == 0 {
		errorObject := types.GraphQLError{
			Message: "unregistered email",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	// validates passwored
	if err := auth.CheckPasswordHash(payload.Input.Password, resp.Data_users[0].Password); err != nil {
		errorObject := types.GraphQLError{
			Message: "wrong password",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	// generates jwt token
	token, err := auth.MakeJWT(resp.Data_users[0].User_id, []byte(os.Getenv("JWT_KEY")), resp.Data_users[0].Role, time.Hour*24)
	if err != nil {
		errorObject := types.GraphQLError{
			Message: "unable login",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

    // set cookie
	http.SetCookie(w, auth.SetCookie(token, 86400))

	// Write the response as JSON
	res := types.LoginResponse{
		Success: true,
	}
	data, _ := json.Marshal(res)
	w.Write(data)

}


