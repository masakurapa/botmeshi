package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/infrastructure"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/usecase"
)

func main() {
	search, err := infrastructure.NewSearchClient()
	if err != nil {
		panic(err.Error())
	}

	notification := infrastructure.NewNotificationClient()
	service := service.NewSearchService(search, notification)
	uc := usecase.NewSearchUseCase(service)
	lambda.Start(handler.NewSearchHandler(uc).Handler)
}
