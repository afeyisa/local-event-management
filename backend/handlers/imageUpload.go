package handlers

import (
	"encoding/json"
	"io"
	models "local-event-management-backend/models/images"
	"local-event-management-backend/types"
	"net/http"
)

func ImageUploads(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	defer r.Body.Close()

	reqBody,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var payload types.ImagesPayload
	err =  json.Unmarshal(reqBody,&payload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	result := models.SaveImages(payload.Input)
	json.NewEncoder(w).Encode(result)

}