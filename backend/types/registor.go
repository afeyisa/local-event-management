package types

type InsertUserArgs struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// to hold user's information which going to be inserted including session variables
type InsertUserPayload struct {
    SessionVariables map[string]interface{} `json:"session_variables"`
    Input            InsertUserArgs         `json:"input"`
}

