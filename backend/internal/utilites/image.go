package utilites

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveImage(
	file multipart.File,
	header *multipart.FileHeader,
	uploadDir string,
) (string, error) {
	if file == nil {
		return "", errors.New("no file provided")
	}

	// genreate unique file name
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(uploadDir, fileName)

	out, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return strings.TrimPrefix(savePath, "static/"), nil
}

func GetImageURL(imagePath string) string {
	if imagePath == "" {
		return ""
	}

	if !strings.HasPrefix(imagePath, "/") {
		imagePath = "/" + imagePath
	}

	return "http://localhost:8080" + imagePath
}
