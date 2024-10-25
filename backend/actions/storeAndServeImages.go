package actions

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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your frontend URL
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse the multipart form, which allows file uploads
	err := r.ParseMultipartForm(10 << 20) // 10MB file limit
	if err != nil {
		fmt.Println(err,"at file")
		http.Error(w, "Unable to process file", http.StatusBadRequest)
		return
	}

	// Get the featured image from the form data
	file, handler, err := r.FormFile("thumbnailimage")
	if err != nil {
		fmt.Println(err,"at tumbna")
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate image type
	if !isImageFile(handler.Filename) {
		log.Printf("Unsupported file type for featured image: %s", handler.Filename)
		http.Error(w, "Unsupported file type for featured image", http.StatusBadRequest)
		return
	}

	// Create a file on the server where the image will be stored
	fileExt := filepath.Ext(handler.Filename) // Get the file extension
	fileBase := strings.TrimSuffix(handler.Filename, fileExt)
	filenameinuuid := uuid.New().String()
	uniqueFileName := fmt.Sprintf("%s_%s%s", fileBase, filenameinuuid, fileExt)

	// Create a new file path with the unique filename
	filePath := filepath.Join(uploadPath, uniqueFileName)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file content to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Handling other images
	var images []string
	files := r.MultipartForm.File["other_images"] // Access all other images

	if len(files) > 4 {
		log.Println("Too many other images uploaded; max is 4.")
		http.Error(w, "Too many other images", http.StatusBadRequest)
		return
	}

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Error opening other image %d: %v", i, err)
			http.Error(w, "Error retrieving other images", http.StatusInternalServerError)
			return
		}

		// Validate image type
		if !isImageFile(fileHeader.Filename) {
			log.Printf("Unsupported file type for other image %d: %s", i, fileHeader.Filename)
			http.Error(w, "Unsupported file type for other images", http.StatusBadRequest)
			return
		}


		fileBase := strings.TrimSuffix(handler.Filename, fileExt)
		filenameinuuid := uuid.New().String()
		uniqueFileName := fmt.Sprintf("%s_%s%s", fileBase, filenameinuuid, fileExt)
		imagePath := filepath.Join(uploadPath, uniqueFileName)

		images = append(images, imagePath)

		// Save each other image
		outFile, err := os.Create(imagePath)
		if err != nil {
			log.Printf("Error saving other image %d: %v", i, err)
			http.Error(w, "Error saving other image", http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(outFile, file); err != nil {
			log.Printf("Error writing other image %d: %v", i, err)
			http.Error(w, "Error saving other image", http.StatusInternalServerError)
			return
		}
		outFile.Close()
		file.Close()
	}

	// Generate the URL to access the uploaded images
	response := map[string]interface{}{
		"thumbnail_image_url": fmt.Sprintf("http://localhost:4000/uploads/%s", uniqueFileName),
		"other_images_urls":  convertToURLs(images),
	}

	// Set response header and send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


func ServeImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image serving called")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your frontend URL
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Extract the image file name from the URL path
	imageName := filepath.Base(r.URL.Path)
	imagePath := filepath.Join(uploadPath, imageName)

	// Check if the file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Check if the file is an image
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

	// Set the appropriate Content-Type header based on the file extension
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

	// Write the image to the response
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error sending the file", http.StatusInternalServerError)
	}
}
