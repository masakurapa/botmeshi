package handler

import (
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
)

// Handler base handler interface
type Handler interface {
	Handler(http.Request) (http.Response, error)
}

// SearchHandler base handler interface
type SearchHandler interface {
	Handler(search.Request) (string, error)
}
