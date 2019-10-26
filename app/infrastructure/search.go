package infrastructure

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/util"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
	"googlemaps.github.io/maps"
)

type searchClient struct {
	maps *maps.Client
	cs   *customsearch.Service
	log  log.Logger
}

// NewSearchClient returns MapSearch instance
func NewSearchClient(logger log.Logger) (repository.Search, error) {
	mc, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		logger.Error("Google Maps Client initialize error", err)
		return nil, fmt.Errorf("map search client initialize error")
	}

	client := &http.Client{Transport: &transport.APIKey{Key: util.CustomSearchAPIKey()}}
	cs, err := customsearch.New(client)
	if err != nil {
		logger.Error("Google Custom Search Client initialize error", err)
		return nil, fmt.Errorf("custom search client initialize error")
	}

	return &searchClient{maps: mc, cs: cs, log: logger}, nil
}

// Station 駅検索
func (c *searchClient) Station(q *search.Query) *search.Station {
	c.log.Start("SearchClient", "Station", q)

	r := &maps.FindPlaceFromTextRequest{
		Input:     q.AreaName + "駅",
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
		Fields:    []maps.PlaceSearchFieldMask{"name", "geometry"},
	}

	c.log.Info("GoogleMaps FindPlaceFromText parameters", r)
	resp, err := c.maps.FindPlaceFromText(context.Background(), r)
	if err != nil {
		c.log.Error("FindPlaceFromText error", err)
		return nil
	}
	if len(resp.Candidates) == 0 {
		c.log.Info("Station not found")
		return nil
	}

	// きっと先頭がその駅のハズだ
	s := &resp.Candidates[0]

	c.log.End("SearchClient", "Station")
	return &search.Station{
		Name: s.Name,
		Location: search.Location{
			Lat: s.Geometry.Location.Lat,
			Lng: s.Geometry.Location.Lng,
		},
	}
}

func (c *searchClient) Shop(q *search.SearchShopsQuery) *search.Shop {
	c.log.Info("Start SearchClient.Shop()", q)

	query := q.Query + " " + q.ShopName
	cx := util.SearchEngineID()
	site, err := c.cs.Cse.Siterestrict.List(query).Cx(cx).Do()
	if err != nil {
		c.log.Error("Shop search error", err)
		return nil
	}

	if len(site.Items) == 0 {
		c.log.Info("Shop not found")
		return nil
	}

	shop := &search.Shop{
		URL: site.Items[0].FormattedUrl,
	}

	c.log.End("SearchClient", "Shop")
	return shop
}

// Shops 検索ワードから店検索
func (c *searchClient) Shops(q *search.SearchShopsQuery) []search.Shop {
	c.log.Info("Start SearchClient.Shops()", q)

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

	c.log.Info("GoogleMaps TextSearch parameters", r)
	resp, err := c.maps.TextSearch(context.Background(), r)
	if err != nil {
		c.log.Error("TextSearch error", err)
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

	c.log.End("SearchClient", "Shops")
	return ret
}
