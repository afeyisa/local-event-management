package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image serving called")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	imageName := filepath.Base(r.URL.Path)
	imagePath := filepath.Join(uploadPath, imageName)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if !isImageFile(imageName) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Open the file
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	contentType := ""
	switch strings.ToLower(filepath.Ext(imageName)) {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	default:
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error sending the file", http.StatusInternalServerError)
	}
}