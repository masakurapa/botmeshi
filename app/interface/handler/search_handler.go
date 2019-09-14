package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type searchHandler struct {
	uc usecase.SearchUseCase
}

// NewSearchHandler returns Handler instance
func NewSearchHandler(uc usecase.SearchUseCase) SearchHandler {
	return &searchHandler{
		uc: uc,
	}
}

// Handler function
func (h *searchHandler) Handler(req search.Request) (string, error) {
	if err := h.uc.Validate(&req); err != nil {
		return err.Error(), nil
	}
	if err := h.uc.Exec(&req); err != nil {
		return err.Error(), nil
	}
	return "Success Search", nil
}
