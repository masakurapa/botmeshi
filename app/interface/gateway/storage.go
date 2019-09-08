package gateway

import (
	"github.com/masakurapa/botmeshi/app/domain/model"
	"github.com/masakurapa/botmeshi/app/domain/repository"
)

// storage struct
type storage struct {
}

// NewStorage returns Storage instance
func NewStorage() repository.Storage {
	return &storage{}
}

// Get file
func (s *storage) Get(path string) (*model.File, error) {
	return model.NewFile("", "", ""), nil
}

// Put file
func (s *storage) Put(file *model.File) error {
	return nil
}
