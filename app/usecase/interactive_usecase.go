package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/util"
)

type interactiveUseCase struct {
	service service.InteractiveService
}

// NewInteractiveUseCase return UseCase instance
func NewInteractiveUseCase(s service.InteractiveService) UseCase {
	return &interactiveUseCase{service: s}
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
	if body.Token != util.APIVerificationToken() {
		return fmt.Errorf("token error")
	}

	return nil
}

// Exec event
func (uc *interactiveUseCase) Exec(body *api.Parameter) error {
	return nil
}
