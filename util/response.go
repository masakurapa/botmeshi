package util

import "github.com/aws/aws-lambda-go/events"

func Response(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "OPTIONS,POST",
			"Access-Control-Allow-Origin":  "*",
		},
		StatusCode: 200,
		Body:       body,
	}
}
