package repository

import "github.com/masakurapa/botmeshi/app/domain/model/search"

// InvokeFunction interface
type InvokeFunction interface {
	Exec(*search.Request) error
}
