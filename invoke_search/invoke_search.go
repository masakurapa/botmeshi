package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/slack-bot/config"
	"github.com/masakurapa/slack-bot/util"
	"github.com/nlopes/slack"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
	"googlemaps.github.io/maps"
)

const (
	shopMax = 5
)

type event struct {
	Channel string `json:"channel"`
	Query   string `json:"query"`
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest func
func HandleRequest(request event) (string, error) {
	s, f := parseQuery(request.Query)
	log.Printf("station: %s, food: %s", s, f)

	// 駅情報
	loc, ok := stationSearch(s)
	if !ok {
		return "OK", nil
	}

	// 店探し
	resp, ok := textSearch(s, f, loc)
	if !ok {
		return "OK", nil
	}

	api := slack.New(util.BotAccessToken())

	if len(resp) == 0 {
		api.PostMessage(request.Channel, slack.MsgOptionText("お店が見つからなかったよ", false))
		return "OK", nil
	}

	text := "お店を見つけたよ！！\n```\n"
	var opts []slack.AttachmentActionOption

	client := &http.Client{Transport: &transport.APIKey{Key: util.CustomSearchAPIKey()}}
	srv, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
		return "OK", nil
	}

	cx := util.SearchEngineID()

	// ランダム5店舗
	shops := random(resp)
	for _, shop := range shops {
		site, err := srv.Cse.Siterestrict.List(request.Query + " " + shop.Name).Cx(cx).Do()
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf("%+v", site.Items[0])

		text += shop.Name + "\n" + site.Items[0].FormattedUrl + "\n"
		opts = append(opts, slack.AttachmentActionOption{
			Text:  shop.Name,
			Value: shop.Name,
		})
	}

	text += "```\n"

	// 店の情報だけまず送る
	api.PostMessage(request.Channel, slack.MsgOptionText(text, false))

	// interactive
	opt := slack.MsgOptionAttachments(slack.Attachment{
		Text:       "いいお店は見つかったかな？",
		CallbackID: "shop",
		Color:      "#ff6633",
		Actions: []slack.AttachmentAction{
			{
				Name:    config.ActionTypeGo,
				Type:    "select",
				Options: opts,
			},
			{
				Name:  config.ActionTypeDoNotGo,
				Type:  "button",
				Text:  "いい店はなかった",
				Style: "danger",
			},
		},
	})
	api.PostMessage(request.Channel, opt)

	return "OK", nil
}

// クエリ文字列を地名・料理名に分割
func parseQuery(query string) (string, string) {
	s := strings.Split(query, " ")
	i := len(s) - 1
	return strings.Join(s[0:i], " "), s[i]
}

// 駅の位置情報を持ってくる
func stationSearch(s string) (*maps.LatLng, bool) {
	c, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	r := &maps.FindPlaceFromTextRequest{
		Input:     s + "駅",
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
		Fields:    []maps.PlaceSearchFieldMask{"name", "geometry"},
	}

	log.Printf("%+v", r)
	resp, err := c.FindPlaceFromText(context.Background(), r)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	log.Printf("%+v", resp.Candidates)

	if len(resp.Candidates) == 0 {
		return nil, true
	}

	// きっと先頭がその駅のハズだ
	return &resp.Candidates[0].Geometry.Location, true
}

func textSearch(s, f string, loc *maps.LatLng) ([]maps.PlacesSearchResult, bool) {
	c, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	r := &maps.TextSearchRequest{
		Query:    s + " " + f,
		Region:   "jp",
		Language: "ja",
		Type:     maps.PlaceTypeRestaurant,
	}

	if loc != nil {
		r.Query = f
		r.Location = loc
		r.Radius = 500
	}

	log.Printf("%+v", r)
	resp, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	log.Printf("%+v", resp.Results)

	return resp.Results, true
}

func random(places []maps.PlacesSearchResult) []maps.PlacesSearchResult {
	if len(places) <= shopMax {
		return places
	}

	rand.Seed(time.Now().UnixNano())

	var num []int
	var ret []maps.PlacesSearchResult
	for len(ret) < shopMax {
		i := rand.Intn(len(places))
		for _, n := range num {
			if n == i {
				continue
			}
		}
		num = append(num, i)
		ret = append(ret, places[i])
	}

	return ret
}
