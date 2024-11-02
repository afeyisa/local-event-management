package actions

import "github.com/google/uuid"

// comman to actions package login and register
type GraphQLError struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Success bool `json:"success"`
}


// login types
type User struct {
	Password string    `json:"password"`
	ID       uuid.UUID `json:"user_id"`
}

type GraphqlResponse struct {
	Data struct {
		User []User `json:"data_users"`
	} `json:"data"`
	Errors []GraphQLError `json:"errors"`
}


type GraphQLRequest struct {
	Query string `json:"query"`
}


type loginArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            loginArgs              `json:"input"`
}


// register types

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


//logout types
type LogOutResponse struct{
	Success bool `json:"success"`
}
