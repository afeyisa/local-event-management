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

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	var payload types.LoginPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return

	}

	resp, err := models.FetchHashedPasswordWithIdByEmail(payload.Input.Email)

	if err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,"failed to get user")

		return
	}

	if len(resp.Data_users) == 0 {
		httpHelper.WriteError(w,http.StatusBadRequest,"unregistered email")

		return
	}

	// validates passwored
	if err := auth.CheckPasswordHash(payload.Input.Password, resp.Data_users[0].Password); err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"wrong password")

		return
	}

	// generates jwt token
	token, err := auth.MakeJWT(resp.Data_users[0].User_id, []byte(os.Getenv("JWT_KEY")), resp.Data_users[0].Role, time.Hour*24)
	if err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,"unable login")
		return
	}

    // set cookie
	http.SetCookie(w, auth.SetCookie(token, 86400))

	res := types.LoginResponse{
		Success: true,
	}
	json.NewEncoder(w).Encode(res)

}


