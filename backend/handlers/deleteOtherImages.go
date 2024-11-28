package handlers

import (
	"encoding/json"
	"io"
	models "local-event-management-backend/models/images"
	httpHelper "local-event-management-backend/helpers/http"
	"local-event-management-backend/types"
	"net/http"
)




func DeleteImages(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	var data types.ImageUrl
	if err := json.Unmarshal(reqBody, &data); err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	path := models.GetPath(data.URL)
	if path == "" {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid image path")
		return
	}

	if err := models.DeleteImage(path); err != nil {
		httpHelper.WriteError(w,http.StatusInternalServerError,"failed to delete image")
		return
	}

	w.WriteHeader(http.StatusOK)
}
