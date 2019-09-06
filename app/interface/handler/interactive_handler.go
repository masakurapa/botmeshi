package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type interactiveHandler struct {
	uc usecase.UseCase
}

// NewInteractiveHandler returns Handler instance
func NewInteractiveHandler(uc usecase.UseCase) Handler {
	return &interactiveHandler{uc: uc}
}

// Handler function
func (h *interactiveHandler) Handler(req http.Request) (http.Response, error) {
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

	return http.NewResponse(http.StatusOK, "Success Interactive"), nil
}
