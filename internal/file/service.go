package file

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	f "go-file-server-api/internal/folder"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Service interface {
	upload(folder string, file multipart.File, header multipart.FileHeader) (string, error)
	read(folder, file string) (string, error)
	delete(folder, file string) error
}

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) upload(folder string, file multipart.File, header multipart.FileHeader) (string, error) {
	uploadDir, err := f.UseUploadDir()
	if err != nil {
		return "", err
	}

	// Checks if folder exists
	// Checks if consumer is in upload dir only
	sanitizedFolder := f.SanitizeName(folder)
	folderPath := filepath.Join(uploadDir, sanitizedFolder)
	if !f.IsInUploadDir(folderPath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return "", errors.New("folder does not exist")
	}

	// Checks if file already exists
	fileName := fmt.Sprintf("%s_%s", uuid.New(), header.Filename)
	filePath := filepath.Join(uploadDir, sanitizedFolder, f.SanitizeName(fileName))
	if !f.IsInUploadDir(filePath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return "", errors.New("file already exists")
	}

	// Creating the output destination of the file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copying the file to local machine
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (s ServiceImpl) read(folder, file string) (string, error) {
	uploadDir, err := f.UseUploadDir()
	if err != nil {
		return "", err
	}

	// Checks if folder exists
	// Checks if consumer is in upload dir only
	sanitizedFolder := f.SanitizeName(folder)
	folderPath := filepath.Join(uploadDir, sanitizedFolder)
	if !f.IsInUploadDir(folderPath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return "", errors.New("folder does not exist")
	}

	// Checks if file already exists
	filePath := filepath.Join(uploadDir, sanitizedFolder, f.SanitizeName(file))
	if !f.IsInUploadDir(filePath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("file doesn't exists")
	}

	return filePath, nil
}

func (s ServiceImpl) delete(folder, file string) error {
	uploadDir, err := f.UseUploadDir()
	if err != nil {
		return err
	}

	// Checks if folder exists
	// Checks if consumer is in upload dir only
	sanitizedFolder := f.SanitizeName(folder)
	folderPath := filepath.Join(uploadDir, sanitizedFolder)
	if !f.IsInUploadDir(folderPath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return errors.New("folder does not exist")
	}

	// Checks if file already exists
	filePath := filepath.Join(uploadDir, sanitizedFolder, f.SanitizeName(file))
	if !f.IsInUploadDir(filePath) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("file doesn't exists")
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
