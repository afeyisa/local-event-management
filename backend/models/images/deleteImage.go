package models

import (
	"os"
	"path/filepath"
	"strings"
)

func DeleteImage(path string) error {
	return os.Remove(path)
}


func GetPath(path string) string{
	baseURL 	:= os.Getenv("BASE_URL")
	uploadPath  := os.Getenv("UPLOAD_PATH") 

	if baseURL == "" || uploadPath == "" {
		return ""
	}
	if !strings.HasPrefix(path, baseURL) {
		return ""
	}

	relativePath := strings.TrimPrefix(path, baseURL)
	absoluteFilePath := filepath.Join(uploadPath, relativePath)
	return absoluteFilePath
}