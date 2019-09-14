package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/usecase"
)

type eventHandler struct {
	uc  usecase.EventUseCase
	log log.Logger
}

// NewEventHandler returns Handler instance
func NewEventHandler(uc usecase.EventUseCase, logger log.Logger) Handler {
	return &eventHandler{uc: uc, log: logger}
}

// Handler function
func (h *eventHandler) Handler(req http.Request) (http.Response, error) {
	h.log.Start("EventHandler", "Handler", req)

	p, err := h.uc.Parse(req.Body)
	if err != nil {
		h.log.Error("Parse error", err)
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	if err = h.uc.Validate(p); err != nil {
		h.log.Error("Validate error", err)
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	if err = h.uc.Exec(p); err != nil {
		h.log.Error("Exec error", err)
		return http.NewResponse(http.StatusOK, err.Error()), nil
	}

	resp := http.NewResponse(http.StatusOK, "Success Event")
	h.log.End("EventHandler", "Handler", resp)
	return resp, nil
}
