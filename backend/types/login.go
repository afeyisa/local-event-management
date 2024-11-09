package types


type LoginResponse struct {
    Success bool `json:"success"`
}

type loginArgs struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginPayload struct {
    SessionVariables map[string]interface{} `json:"session_variables"`
    Input            loginArgs              `json:"input"`
}
