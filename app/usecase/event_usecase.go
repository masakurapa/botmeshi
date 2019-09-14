package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/util"
)

// EventUseCase interface
type EventUseCase interface {
	Parse(string) (*api.Parameter, error)
	Validate(*api.Parameter) error
	Exec(*api.Parameter) error
}

type eventUseCase struct {
	service service.EventService
}

// NewEventUseCase return EventUseCase instance
func NewEventUseCase(s service.EventService) EventUseCase {
	return &eventUseCase{service: s}
}

// Parse request
func (uc *eventUseCase) Parse(body string) (*api.Parameter, error) {
	eb := api.EventBody{}
	if err := json.Unmarshal([]byte(body), &eb); err != nil {
		return nil, err
	}
	return eb.ToParameter(), nil
}

// Validate request
func (uc *eventUseCase) Validate(body *api.Parameter) error {
	// check token
	if body.Token != util.APIVerificationToken() {
		return fmt.Errorf("token error")
	}

	// Event APIの認証用
	if body.Type == "url_verification" {
		return fmt.Errorf(body.Challenge)
	}

	return nil
}

// Exec event
func (uc *eventUseCase) Exec(body *api.Parameter) error {
	return uc.service.Exec(body)
}
