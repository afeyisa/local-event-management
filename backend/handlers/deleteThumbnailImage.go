package handlers

import (
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	models "local-event-management-backend/models/images"
	"local-event-management-backend/types"
	"net/http"
)




func DeleteThumbnail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	var data types.ThumbNailUrl
	if err := json.Unmarshal(reqBody, &data); err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
		return
	}

	path := models.GetPath(data.Thumbnail_image_url)
	if path == "" {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid thumbnail path")
		return
	}

	if err := models.DeleteImage(path); err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"failed to delete thumbnail")
		return
	}

	w.WriteHeader(http.StatusOK)
}



