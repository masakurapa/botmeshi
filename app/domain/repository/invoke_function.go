package repository

import "github.com/masakurapa/botmeshi/app/domain/model/invoke"

// InvokeFunction interface
type InvokeFunction interface {
	Exec(*invoke.Parameter) error
}
