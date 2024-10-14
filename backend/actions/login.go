package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/auth"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type User struct {
    Password string `json:"password"`
	ID uuid.UUID `json:"user_id"`
}

type HasuraUsersResponse struct {
    Data struct {
        User []User`json:"data_users"`
    } `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

type GraphQLRequest struct {
	Query     string         `json:"query"`
}

// type GraphQLError struct {
// 	Message string `json:"message"`
// }
type LoginResponse struct {
	Success bool `json:"success"`
}

type loginArgs struct {
	Email   string `json:"email"`
	Password string    `json:"password"`
}

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            loginArgs              `json:"input"`
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// parse the body as action payload
	var actionPayload ActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		errorObject := GraphQLError{
			Message: "invalid payload",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
		
	}

	
	resp, err := fetchHashedPasswordWithId(actionPayload.Input.Email)

	if err != nil {
		errorObject := GraphQLError{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	if len(resp.Errors) > 0 {
		errorObject := GraphQLError{
			Message: "unable to access user data",
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	if len(resp.Data.User) == 0{
		errorObject := GraphQLError{
			Message: "unregistered email",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	// validates passwored
	if err := auth.CheckPasswordHash(actionPayload.Input.Password, resp.Data.User[0].Password); err != nil {
		errorObject := GraphQLError{
			Message: "wrong password",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}

	// generates jwt token
	token, err := auth.MakeJWT(resp.Data.User[0].ID, []byte(os.Getenv("JWT_KEY")), time.Hour*24)
	if err != nil {
		errorObject := GraphQLError{
			Message: "unable login",
		}
		w.WriteHeader(http.StatusBadRequest)
		errorBody, _ := json.Marshal(errorObject)
		w.Write(errorBody)
		return
	}
	
	// set cookie

	cookie := http.Cookie{
	Name:     "jwttoken",
	Value:   token,
	Path:     "/",
	HttpOnly: true,
	Secure:   true,
	SameSite: http.SameSiteDefaultMode,
	MaxAge:   86400, // Cookie expiration (1 day)
    }

    http.SetCookie(w, &cookie)

	// Write the response as JSON
	res := LoginResponse{
		Success: true,
	}
	data, _ := json.Marshal(res)
	w.Write(data)

}


// Function to fetch hashed password from Hasura for the given user ID
func fetchHashedPasswordWithId(email string) (HasuraUsersResponse, error) {
	query := fmt.Sprintf(`{
		hasura_users(where: {email: {_eq: "%s"}}) {
		  password
		  user_id
		}
	  }`, email)

	var hasuraResp HasuraUsersResponse
	reqBody := GraphQLRequest{
		Query:query ,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return hasuraResp, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/graphql", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return hasuraResp, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return hasuraResp, fmt.Errorf("failed to request user data")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return hasuraResp, fmt.Errorf("failed to fetch user data")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return hasuraResp, fmt.Errorf("failed to read user data")
	}

	// bad request
	if err := json.Unmarshal(respBody, &hasuraResp); err != nil {
		return hasuraResp, fmt.Errorf("failed to unmarshal user data")
	}
	return hasuraResp, nil
}

