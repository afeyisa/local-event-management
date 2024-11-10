package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/models/images"
	"local-event-management-backend/types"
	"net/http"
)

func ImageUploads(w http.ResponseWriter, r *http.Request){
	fmt.Println("saving the images")
	w.Header().Set("Content-Type","application/json")

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

	result := images.SaveImages(payload.Input)
	data, err := json.Marshal(result)

	if err != nil{
		errorObject := types.GraphQLError{
			Message: " un able to create links",
		}
	err ,_ := json.Marshal(errorObject)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}