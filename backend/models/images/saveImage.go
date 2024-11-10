package models

import (
	"encoding/base64"
	"fmt"
	"local-event-management-backend/types"
	"os"
	"strings"

	"github.com/google/uuid"
)

var uploadPath = "./uploads/"

func isImageFile(ext string) bool {
	return ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif"
}

func saveHelper(image string)(url string,err error){

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
	url = fmt.Sprintf("http://localhost:4000/uploads/%s.%s", fileName, extension)
	return url, nil
}


func SaveImages(payload types.Images)(res types.Urls){
	fmt.Println("saving the from the model")
	var errors []string
	if(payload.ThumbNail !=""){
		url, saveErr := saveHelper(payload.ThumbNail)
		fmt.Println(url," test1")
		if saveErr != nil {
			errors = append(errors, saveErr.Error())
		} else{
			fmt.Println(url)
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