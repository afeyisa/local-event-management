package handlers

import (
	"encoding/json"
	"io"
	models "local-event-management-backend/models/images"
	"local-event-management-backend/types"
	"net/http"
)



func ServeImage(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")

  reqBody, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  var payload types.ImageRequestPayload
  err = json.Unmarshal(reqBody, &payload)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  result, err := models.GetImage(payload.Input)

  if err != nil {
    errorObject := types.GraphQLError{
      Message: err.Error(),
    }
    errorBody, _ := json.Marshal(errorObject)
    w.WriteHeader(http.StatusBadRequest)
    w.Write(errorBody)
    return
  }

  data, _ := json.Marshal(result)
  w.Write(data)

}
