package main

import (
	"context"
	"encoding/json"
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	l "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/masakurapa/slack-bot/config"
	"github.com/masakurapa/slack-bot/util"
)

type payload struct {
	Type    string   `json:"type"`
	Token   string   `json:"token"`
	Channel channel  `json:"channel"`
	Actions []action `json:"actions"`
}

type channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type action struct {
	Name            string           `json:"name"`
	Value           string           `json:"value"`
	SelectedOptions []selectedOption `json:"selected_options"`
}

type selectedOption struct {
	Value string `json:"value"`
}

type invokeParam struct {
	Channel string `json:"channel"`
	Query   string `json:"query"`
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest func
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("%+v", request)

	// 先頭8文字(payload=)以外を使う
	jsonStr, err := url.QueryUnescape(request.Body[8:])
	if err != nil {
		log.Fatal(err)
		return util.Response("エラーが発生した!!"), nil
	}

	var body payload
	if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
		log.Fatal(err)
		return util.Response("エラーが発生した!!"), nil
	}

	log.Printf("%+v", body)

	// check token
	if body.Token != util.BotVerificationToken() {
		log.Fatalln("invalid token: " + body.Token)
		return util.Response("キサマ何者だ！"), nil
	}

	var msg string
	switch body.Actions[0].Name {
	case config.ActionTypeCancel:
		msg = "ばいびー"
	case config.ActionTypeGo:
		msg = "`" + body.Actions[0].SelectedOptions[0].Value + "` に\nごーーーーーーーーーーーーーーーる！！"
	case config.ActionTypeDoNotGo:
		msg = "ざんねん...。"
	case config.ActionTypeSelect:
		msg = invoke(body.Channel.ID, body.Actions[0].SelectedOptions[0].Value)
	default:
		msg = "キサマ何者だ！"
	}

	return util.Response(msg), nil
}

func invoke(channel, query string) string {
	s, err := json.Marshal(invokeParam{
		Channel: channel,
		Query:   query,
	})
	if err != nil {
		log.Fatal(err)
		return "エラーが発生した!!"
	}

	lmd := l.New(session.New())
	_, err = lmd.Invoke(&l.InvokeInput{
		FunctionName:   aws.String(util.InvokeLambdaArn()),
		Payload:        []byte(s),
		InvocationType: aws.String("Event"),
	})

	if err != nil {
		log.Fatal(err)
		return "エラーが発生した!!"
	}

	return "`" + query + "` でお店を探すよ！\nちょっと時間がかかるからまってくれ！"
}
