package models

import (
	"encoding/base64"
	"fmt"
	"local-event-management-backend/types"
	"os"
	"strings"

	"github.com/google/uuid"
)


func isImageFile(ext string) bool {
	return ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif"|| ext == "webp"
}

func saveHelper(image string)(url string,err error){
	var uploadPath = os.Getenv("UPLOAD_PATH")
	fmt.Println("upload path in saving",uploadPath)
	parts := strings.SplitN(image,",",2)

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid base64 image format")
	}

	metadata := parts[0]
	imageData := parts[1]

	metadataParts := strings.Split(metadata, "/")
	if len(metadataParts) < 2 {
		return "", fmt.Errorf("invalid metadata format")
	}

	extensionParts := strings.Split(metadataParts[1], ";")
	extension := extensionParts[0]
	if !isImageFile(extension){
		return "", fmt.Errorf("unsuported file type")
	}

	dst, decodeErr := base64.StdEncoding.DecodeString(imageData)

    if decodeErr != nil {
		return "", fmt.Errorf("failed to decode base64 image: %v", decodeErr)
    }

	fileName := uuid.New().String()

    writeErr := os.WriteFile(fmt.Sprintf("%v%v.%v",uploadPath,fileName,extension), dst, 0666)
    if writeErr != nil {
		return "", fmt.Errorf("failed to write image file: %v", writeErr)
    }
	url = fmt.Sprintf("%s/%s.%s", os.Getenv("BASE_URL"), fileName, extension)
	return url, nil
}


func SaveImages(payload types.Images)(res types.Urls){
	var errors []string
	if(payload.ThumbNail !=""){
		url, saveErr := saveHelper(payload.ThumbNail)
		if saveErr != nil {
			errors = append(errors, saveErr.Error())
		} else{
			res.Thumbnail_image_url = url}
	}
	if len(payload.OtherImages) > 0 {
		var otherUrls []string
		for _, image := range payload.OtherImages{

			if image != "" {
				url, saveErr := saveHelper(image)
				if saveErr != nil {
					errors = append(errors, saveErr.Error())
				}else {
				otherUrls = append(otherUrls, url)
			}
			}
		}
		res.Other_images_urls = otherUrls

	}
	res.Errors = errors

	return res
}