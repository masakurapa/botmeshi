package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
)

// Handler base handler interface
type Handler interface {
	Handler(req http.Request) (http.Response, error)
}

// SearchHandler base handler interface
type SearchHandler interface {
	Handler(req search.Request) (string, error)
}
