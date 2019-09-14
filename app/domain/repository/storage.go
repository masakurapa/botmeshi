package repository

import "github.com/masakurapa/botmeshi/app/domain/model"

// Storage interface
type Storage interface {
	Get(path string) (*model.File, error)
	Put(file *model.File) error
}
