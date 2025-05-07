package folder

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Service interface {
	create(name string) error // create outputs home + uploadDir + name
	remove(name string) error // remove folder inside uploadDir recursively
}

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) create(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is empty")
	}

	uploadDir, err := UseUploadDir()
	if err != nil {
		return err
	}

	path := filepath.Join(uploadDir, filepath.Clean(name))
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return errors.New("folder already exists")
	}

	return nil
}

func (s ServiceImpl) remove(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is empty")
	}

	uploadDir, err := UseUploadDir()
	if err != nil {
		return err
	}

	path := filepath.Join(uploadDir, name)
	err = os.RemoveAll(path)
	if err != nil {
		return errors.New("folder not exists")
	}

	return nil
}
