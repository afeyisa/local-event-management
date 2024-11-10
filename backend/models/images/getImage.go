package models

import (
	"encoding/base64"
	"fmt"
	"local-event-management-backend/types"
	"os"
	"path/filepath"
	"strings"
)

var path = "./uploads"

func GetImage(link types.Url) (response types.Image, err error) {

	const baseURL = "http://localhost:4000/uploads"

	if !strings.HasPrefix(link.Link, baseURL) {
		return response, fmt.Errorf("invalid link format")
	}

	relativePath := strings.TrimPrefix(link.Link, baseURL)
	absoluteFilePath := filepath.Join( path, relativePath)
	extension := strings.TrimPrefix(filepath.Ext(absoluteFilePath), ".")


	if !isImageFile(extension){
		return response, fmt.Errorf("invalid file format")
	}
	
	data, err := os.ReadFile(absoluteFilePath)
	if err != nil {
		return response, fmt.Errorf("error reading file")
    }

	str := base64.StdEncoding.EncodeToString(data)
    dataURI := fmt.Sprintf("data:image/%v;base64,%s",extension, str)   
	return types.Image{Image: dataURI}, nil
}