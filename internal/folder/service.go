package folder

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Service interface {
	create(name string) (sanitizeName string, err error) // create outputs home + uploadDir + name. sanitizeName returns the saved folder name
	remove(name string) error                            // remove folder inside uploadDir recursively
}

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) create(name string) (sanitizeName string, err error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name is required")
	}

	uploadDir, err := UseUploadDir()
	if err != nil {
		return "", err
	}

	sanitizeName = SanitizeName(name)

	path := filepath.Join(uploadDir, sanitizeName)
	if !IsInUploadDir(path) {
		panic("error user is not in upload directory. Terminating the program")
	}

	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return "", errors.New("folder already exists")
	}

	return sanitizeName, nil
}

func (s ServiceImpl) remove(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	uploadDir, err := UseUploadDir()
	if err != nil {
		return err
	}

	path := filepath.Join(uploadDir, SanitizeName(name))
	if !IsInUploadDir(path) {
		panic("error user is not in upload directory. Terminating the program")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("folder does not exist")
	}

	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
