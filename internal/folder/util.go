package folder

import (
	"os"
	"path/filepath"
)

// UseUploadDir returns home + uploadDir
func UseUploadDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, os.Getenv("UPLOAD_ROOT_FOLDER"))
	return path, nil
}
