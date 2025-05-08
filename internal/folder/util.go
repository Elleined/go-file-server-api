package folder

import (
	"os"
	"path/filepath"
	"strings"
)

// UseUploadDir returns home + uploadDir
func UseUploadDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, SanitizeName(os.Getenv("UPLOAD_ROOT_FOLDER")))
	return path, nil
}

// SanitizeName name to ensure traversal attacks are prevented. Only use inside the join filepath method
func SanitizeName(name string) string {
	return filepath.Base(filepath.Clean(name))
}

// IsInUploadDir returns whether the user is still in upload dir. Only use after joining path
func IsInUploadDir(path string) bool {
	uploadDir, err := UseUploadDir()
	if err != nil {
		return false
	}

	absUploadDir, err := filepath.Abs(uploadDir)
	if err != nil {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	return strings.HasPrefix(absPath, absUploadDir)
}
