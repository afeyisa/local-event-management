package httpHelper

import (
	"encoding/json"
	"local-event-management-backend/types"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.GraphQLError{Message: message})
}