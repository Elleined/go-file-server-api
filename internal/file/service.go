package file

import (
	"errors"
	"fmt"
	f "go-file-server-api/internal/folder"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Service interface {
	upload(folder string, file multipart.File, header multipart.FileHeader) error
	read(folder, file string) error
	delete(folder, file string) error
}

type ServiceImpl struct {
}

func (s ServiceImpl) upload(folder string, file multipart.File, header multipart.FileHeader) error {
	if strings.TrimSpace(folder) == "" {
		return errors.New("folder is required")
	}

	if file == nil {
		return errors.New("file is required")
	}

	uploadDir, err := f.UseUploadDir()
	if err != nil {
		return err
	}

	// Creating the output destination of the file
	dst, err := os.Create(filepath.Join(uploadDir, folder, filepath.Base(header.Filename)))
	if err != nil {
		return err
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
		return err
	}

	return nil
}

func (s ServiceImpl) read(folder, file string) error {
	fmt.Println("Reading file " + file)
	fmt.Println("Folder", folder)
	return nil
}

func (s ServiceImpl) delete(folder, file string) error {
	fmt.Println("Deleting file " + file)
	fmt.Println("Folder", folder)
	return nil
}

func NewService() Service {
	return &ServiceImpl{}
}
