package repository

import "github.com/masakurapa/botmeshi/app/domain/model/search"

// Search interface
type Search interface {
	Station(*search.Query) *search.Station
	Shop(*search.SearchShopsQuery) *search.Shop
	Shops(*search.SearchShopsQuery) []search.Shop
}
