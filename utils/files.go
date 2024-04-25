package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	maxFileSize     = 5 * 1024 * 1024 // 5 MB
	allowedFileType = "application/pdf"
)

func GenerateFileName(fileName string) string {

	// Generate a unique prefix using UUID
	prefix := uuid.New().String()
	fileExt := filepath.Ext(fileName)
	newFileName := prefix + fileExt

	return newFileName
}

func IsValidFile(file *multipart.FileHeader) bool {
	// Check file size
	if file.Size > maxFileSize {
		return false
	}

	// Check file type
	fileType := file.Header.Get("Content-Type")
	fmt.Println(fileType)
	if !isAllowedFileType(fileType) {
		return false
	}

	return true
}

func isAllowedFileType(fileType string) bool {
	return strings.EqualFold(fileType, allowedFileType)
}
