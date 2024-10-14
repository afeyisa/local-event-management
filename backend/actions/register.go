package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"local-event-management-backend/auth"
	"net/http"
	"os"
	"time"
	"github.com/google/uuid"
)



type RegisterOutput struct {
	User_id uuid.UUID `json:"user_id"`
}
type insertUserArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InsertActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            insertUserArgs         `json:"input"`
}

type InsertGraphQLRequest struct {
	Query     string         `json:"query"`
	Variables insertUserArgs `json:"variables"`
}

type GraphQLData struct {
	Insert_data_users_one RegisterOutput `json:"insert_data_users_one"`
}

type GraphQLResponse struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

func ResgisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test register")
	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// read request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// parse the body as action payload
	var actionPayload InsertActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	// Send the request params to the Action's generated handler function
	result, err := insertUser(actionPayload.Input)
	// throw if an error happens
	if err != nil {
		errorObject := GraphQLError{
			Message: "un able to create user with " + actionPayload.Input.Email,
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}
    token, err := auth.MakeJWT(result.User_id, []byte(os.Getenv("JWT_KEY")), time.Hour*24)
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

// Auto-generated function that takes the Action parameters and must return it's response type
func insertUser(args insertUserArgs) (response RegisterOutput, err error) {
	p, err := auth.HashPassword(args.Password)
	if err != nil{
		return
	}
	args.Password = p
	hasuraResponse, err := execute(args)
	// throw if any unexpected error happens
	if err != nil {
		return
	}
	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}
	response = hasuraResponse.Data.Insert_data_users_one
	return
}

// insert the user and return user_id
func execute(variables insertUserArgs) (response GraphQLResponse, err error) {
	reqBody := InsertGraphQLRequest{
		Query:     "mutation  register($email: String, $password: String) {   insert_data_users_one(object: {email: $email ,password: $password}){  user_id   } }",
		Variables: variables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/graphql", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 
	}
	err = json.Unmarshal(respBody, &response);
	// bad request
	if  err != nil {
		return
	}
	return 
}