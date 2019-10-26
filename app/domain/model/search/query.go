package search

// Query struct
type Query struct {
	AreaName string
	Genre    string
}

type SearchShopsQuery struct {
	Query    string
	ShopName string
	Radius   uint
	Location *Location
}

type Location struct {
	Lat float64
	Lng float64
}
