package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
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
	resp, ok := textSearch(request.Query)
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

		log.Printf("%+v\n", site)

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

func textSearch(query string) ([]maps.PlacesSearchResult, bool) {
	c, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	r := &maps.TextSearchRequest{
		Query:    query,
		Region:   "jp",
		Language: "ja",
		Type:     maps.PlaceTypeRestaurant,
	}

	resp, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	log.Printf("%+v\n", resp.Results)
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
