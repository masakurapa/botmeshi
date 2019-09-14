package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/util"
)

// InteractiveUseCase interface
type InteractiveUseCase interface {
	Parse(string) (*api.Parameter, error)
	Validate(*api.Parameter) error
	Exec(*api.Parameter) (string, error)
}

type interactiveUseCase struct {
	service service.InteractiveService
	log     log.Logger
}

// NewInteractiveUseCase return InteractiveUseCase instance
func NewInteractiveUseCase(s service.InteractiveService, logger log.Logger) InteractiveUseCase {
	return &interactiveUseCase{service: s, log: logger}
}

// Parse request
func (uc *interactiveUseCase) Parse(body string) (*api.Parameter, error) {
	eb := api.InteractiveBody{}
	if err := json.Unmarshal([]byte(body), &eb); err != nil {
		return nil, err
	}
	return eb.ToParameter(), nil
}

// Validate request
func (uc *interactiveUseCase) Validate(body *api.Parameter) error {
	// check token
	if body.Token != util.BotVerificationToken() {
		return fmt.Errorf("token error")
	}

	return nil
}

// Exec event
func (uc *interactiveUseCase) Exec(body *api.Parameter) (string, error) {
	return uc.service.Exec(body)
}
