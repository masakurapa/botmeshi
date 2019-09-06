package handler

import "github.com/masakurapa/botmeshi/app/domain/model/http"

// Handler base handler interface
type Handler interface {
	Handler(req http.Request) (http.Response, error)
}
