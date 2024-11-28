package handlers

import (
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	models "local-event-management-backend/models/images"
	"local-event-management-backend/types"
	"net/http"
)



func ServeImage(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")

  reqBody, err := io.ReadAll(r.Body)
  defer r.Body.Close()
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
    httpHelper.WriteError(w,http.StatusBadRequest,"couldn;t find image")
  }
  json.NewEncoder(w).Encode(result)
}
