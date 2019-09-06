package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type eventHandler struct {
	uc usecase.UseCase
}

// NewEventHandler returns Handler instance
func NewEventHandler(uc usecase.UseCase) Handler {
	return &eventHandler{uc: uc}
}

// Handler function
func (h *eventHandler) Handler(req http.Request) (http.Response, error) {
	p, err := h.uc.Parse(req.Body)
	if err != nil {
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	if err = h.uc.Validate(p); err != nil {
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	if err = h.uc.Exec(p); err != nil {
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	return http.NewResponse(http.StatusOK, "Success Event"), nil
}
