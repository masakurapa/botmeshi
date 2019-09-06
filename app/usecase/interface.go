package usecase

import (
	"github.com/masakurapa/botmeshi/app/domain/model/api"
)

// UseCase interface
type UseCase interface {
	Parse(string) (*api.Parameter, error)
	Validate(*api.Parameter) error
	Exec(*api.Parameter) error
}
