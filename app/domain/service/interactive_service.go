package service

import "github.com/masakurapa/botmeshi/app/domain/repository"

// InteractiveService interface
type InteractiveService interface {
}

type interactiveService struct {
	storage repository.Storage
}

// NewInteractiveService returns InteractiveService instance
func NewInteractiveService(s repository.Storage) InteractiveService {
	return &interactiveService{storage: s}
}
