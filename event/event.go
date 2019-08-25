package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/slack-bot/config"
	"github.com/masakurapa/slack-bot/util"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

type eventBody struct {
	Token     string     `json:"token"`
	Type      string     `json:"type"`
	Event     innerEvent `json:"event"`
	Challenge string     `json:"challenge"`
}
type innerEvent struct {
	Type    string `json:"type"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
	User    string `json:"user"`
}

var menus = []string{
	"ラーメン", "肉", "魚", "定食", "カレー", "和食", "中華",
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest func
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("%+v\n", request)

	var event eventBody
	err := json.Unmarshal([]byte(request.Body), &event)
	if err != nil {
		return util.Response("OK"), nil
	}

	// check token
	if event.Token != util.BotVerificationToken() {
		log.Println("invalid token: " + event.Token)
		return util.Response("OK"), nil
	}

	// Event APIの認証用
	if event.Type == slackevents.URLVerification {
		log.Println("auth event api.")
		return util.Response(event.Challenge), nil
	}

	// mention only
	if event.Event.Type != slackevents.AppMention {
		log.Println("invalid event type: " + event.Type)
		return util.Response("OK"), nil
	}

	api := slack.New(util.BotAccessToken())

	// 先頭12文字はメンション用の固定文字なのでいらない
	text := strings.TrimSpace(event.Event.Text[12:])

	if text == "" {
		api.PostMessage(event.Event.Channel, slack.MsgOptionText("探したい地名をいれてくれ", false))
		return util.Response("OK"), nil
	}

	var opts []slack.AttachmentActionOption
	for i := 0; i < len(menus); i++ {
		opts = append(opts, slack.AttachmentActionOption{
			Text:  menus[i],
			Value: text + " " + menus[i],
		})
	}

	opt := slack.MsgOptionAttachments(slack.Attachment{
		Text:       text + " で何が食べたい？",
		CallbackID: "menu",
		Actions: []slack.AttachmentAction{
			{
				Name:    config.ActionTypeSelect,
				Type:    "select",
				Options: opts,
			},
			{
				Name:  config.ActionTypeCancel,
				Type:  "button",
				Text:  "やめる",
				Style: "danger",
			},
		},
	})

	api.PostMessage(event.Event.Channel, opt)
	return util.Response("OK"), nil
}
