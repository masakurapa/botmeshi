package infrastructure

import (
	"context"

	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/util"
	"googlemaps.github.io/maps"
)

type searchClient struct {
	client *maps.Client
	log    log.Logger
}

// NewSearchClient returns MapSearch instance
func NewSearchClient(logger log.Logger) (repository.Search, error) {
	c, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		return nil, err
	}
	return &searchClient{client: c, log: logger}, nil
}

// Station 駅検索
func (c *searchClient) Station(q *search.Query) *search.Station {
	r := &maps.FindPlaceFromTextRequest{
		Input:     q.AreaName + "駅",
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
		Fields:    []maps.PlaceSearchFieldMask{"name", "geometry"},
	}

	resp, err := c.client.FindPlaceFromText(context.Background(), r)
	if err != nil {
		return nil
	}
	if len(resp.Candidates) == 0 {
		return nil
	}

	// きっと先頭がその駅のハズだ
	s := &resp.Candidates[0]

	return &search.Station{
		Name: s.Name,
		Location: search.Location{
			Lat: s.Geometry.Location.Lat,
			Lng: s.Geometry.Location.Lng,
		},
	}
}

func (c *searchClient) Shop(*search.SearchShopsQuery) *search.Shop {
	return &search.Shop{}
}

// Shops 検索ワードから店検索
func (c *searchClient) Shops(q *search.SearchShopsQuery) []search.Shop {
	r := &maps.TextSearchRequest{
		Query:    q.Query,
		Region:   "jp",
		Language: "ja",
		Type:     maps.PlaceTypeRestaurant,
	}

	if q.Location != nil {
		r.Location = &maps.LatLng{
			Lat: q.Location.Lat,
			Lng: q.Location.Lng,
		}
		r.Radius = q.Radius
	}

	resp, err := c.client.TextSearch(context.Background(), r)
	if err != nil {
		return nil
	}

	var ret []search.Shop
	for _, r := range resp.Results {
		ret = append(ret, search.Shop{
			Name: r.Name,
			Location: search.Location{
				Lat: r.Geometry.Location.Lat,
				Lng: r.Geometry.Location.Lng,
			},
		})
	}

	return ret
}
