package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var uploadPath = "./uploads/"

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func convertToURLs(images []string) []string {
	var urls []string
	for _, image := range images {
		urls = append(urls, "http://localhost:4000/uploads/"+filepath.Base(image))
	}
	return urls
} 

func UploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image saving called")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err, "at file")
		http.Error(w, "Unable to process file", http.StatusBadRequest)
		return
	}

	_, _, thumbnailErr := r.FormFile("thumbnailimage")
	otherImages := r.MultipartForm.File["other_images"]
	if thumbnailErr != nil && len(otherImages) == 0 {
		http.Error(w, "At least one image is required", http.StatusBadRequest)
		return
	}

	var thumbnailURL string
	if thumbnailErr == nil {
		file, handler, err := r.FormFile("thumbnailimage")
		if err == nil {
			defer file.Close()

			if !isImageFile(handler.Filename) {
				log.Printf("Unsupported file type for featured image: %s", handler.Filename)
				http.Error(w, "Unsupported file type for featured image", http.StatusBadRequest)
				return
			}

			fileExt := filepath.Ext(handler.Filename)
			fileBase := strings.TrimSuffix(handler.Filename, fileExt)
			filenameInUUID := uuid.New().String()
			uniqueFileName := fmt.Sprintf("%s_%s%s", fileBase, filenameInUUID, fileExt)

			filePath := filepath.Join(uploadPath, uniqueFileName)
			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Unable to save the file", http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, "Error saving the file", http.StatusInternalServerError)
				return
			}
			thumbnailURL = fmt.Sprintf("http://localhost:4000/uploads/%s", uniqueFileName)
		}
	}

	var images []string
	if len(otherImages) > 0 {
		if len(otherImages) > 4 {
			log.Println("Too many other images uploaded; max is 4.")
			http.Error(w, "Too many other images", http.StatusBadRequest)
			return
		}

		for i, fileHeader := range otherImages {
			file, err := fileHeader.Open()
			if err != nil {
				log.Printf("Error opening other image %d: %v", i, err)
				http.Error(w, "Error retrieving other images", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			if !isImageFile(fileHeader.Filename) {
				log.Printf("Unsupported file type for other image %d: %s", i, fileHeader.Filename)
				http.Error(w, "Unsupported file type for other images", http.StatusBadRequest)
				return
			}

			fileExt := filepath.Ext(fileHeader.Filename)
			fileBase := strings.TrimSuffix(fileHeader.Filename, fileExt)
			filenameInUUID := uuid.New().String()
			uniqueFileName := fmt.Sprintf("%s_%s%s", fileBase, filenameInUUID, fileExt)
			imagePath := filepath.Join(uploadPath, uniqueFileName)
			images = append(images, imagePath)

			outFile, err := os.Create(imagePath)
			if err != nil {
				log.Printf("Error saving other image %d: %v", i, err)
				http.Error(w, "Error saving other image", http.StatusInternalServerError)
				return
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, file); err != nil {
				log.Printf("Error writing other image %d: %v", i, err)
				http.Error(w, "Error saving other image", http.StatusInternalServerError)
				return
			}
		}
	}

	response := map[string]interface{}{
		"thumbnail_image_url": thumbnailURL,
		"other_images_urls":   convertToURLs(images),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}