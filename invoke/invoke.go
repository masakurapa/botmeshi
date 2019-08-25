package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/slack-bot/config"
	"github.com/masakurapa/slack-bot/util"
	"github.com/nlopes/slack"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
	"googlemaps.github.io/maps"
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
	resp, err := textSearch(request.Query)
	if err != nil {
		return "OK", nil
	}

	text := "お店を見つけたよ！！\n```\n"
	var opts []slack.AttachmentActionOption

	client := &http.Client{Transport: &transport.APIKey{Key: util.CustomSearchAPIKey()}}
	srv, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	cx := util.SearchEngineID()

	// ランダム5店舗
	for i := 0; i < 5; i++ {
		site, err := srv.Cse.Siterestrict.List(request.Query + " " + resp[i].Name).Cx(cx).Do()
		if err != nil {
			log.Fatal(err)
			continue
		}

		text += resp[i].Name + "\n" + site.Items[0].FormattedUrl + "\n"
		opts = append(opts, slack.AttachmentActionOption{
			Text:  resp[i].Name,
			Value: resp[i].Name,
		})
	}

	api := slack.New(util.BotAccessToken())

	// 店の情報だけまず送る
	api.PostMessage(request.Channel, slack.MsgOptionText(text, false))

	// interactive
	opt := slack.MsgOptionAttachments(slack.Attachment{
		Text:       text,
		CallbackID: "shop",
		Actions: []slack.AttachmentAction{
			{
				Name:    config.ActionTypeSelect,
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

func textSearch(query string) ([]maps.PlacesSearchResult, error) {
	c, err := maps.NewClient(maps.WithAPIKey(util.PlaceAPIKey()))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, err
	}

	r := &maps.TextSearchRequest{
		Query:    query,
		Region:   "jp",
		Language: "ja",
		Type:     maps.PlaceTypeRestaurant,
	}

	resp, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, err
	}

	return resp.Results, nil
}
