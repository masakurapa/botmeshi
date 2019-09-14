package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type searchHandler struct {
	uc  usecase.SearchUseCase
	log log.Logger
}

// NewSearchHandler returns Handler instance
func NewSearchHandler(uc usecase.SearchUseCase, logger log.Logger) SearchHandler {
	return &searchHandler{uc: uc, log: logger}
}

// Handler function
func (h *searchHandler) Handler(req search.Request) (string, error) {
	h.log.Info("Start SearchHandler: %v", req)

	if err := h.uc.Validate(&req); err != nil {
		h.log.Error("Validate error: %s", err.Error())
		return err.Error(), nil
	}
	if err := h.uc.Exec(&req); err != nil {
		h.log.Error("Exec error: %s", err.Error())
		return err.Error(), nil
	}

	h.log.Info("End SearchHandler")
	return "Success Search", nil
}
