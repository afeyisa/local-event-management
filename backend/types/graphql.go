package types

type GraphQLError struct {
    Message string `json:"message"`
}

type GraphQLRequest struct {
    Query     string      `json:"query"`
    Variables interface{} `json:"variables,omitempty"`
}

