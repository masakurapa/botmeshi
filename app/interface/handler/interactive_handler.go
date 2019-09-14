package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type interactiveHandler struct {
	uc  usecase.InteractiveUseCase
	log log.Logger
}

// NewInteractiveHandler returns Handler instance
func NewInteractiveHandler(uc usecase.InteractiveUseCase, logger log.Logger) Handler {
	return &interactiveHandler{uc: uc, log: logger}
}

// Handler function
func (h *interactiveHandler) Handler(req http.Request) (http.Response, error) {
	h.log.Info("Start InteractiveHandler: %s", req.Body)

	p, err := h.uc.Parse(req.Body)
	if err != nil {
		h.log.Error("Parse error: %s", err.Error())
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	h.log.Info("Parsed parameters: %+v", p)

	if err = h.uc.Validate(p); err != nil {
		h.log.Error("Validate error: %s", err.Error())
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	msg, err := h.uc.Exec(p)
	if err != nil {
		h.log.Error("Exec error: %s", err.Error())
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	h.log.Info("End InteractiveHandler")
	return http.NewResponse(http.StatusOK, msg), nil
}
