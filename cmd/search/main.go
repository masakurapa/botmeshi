package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/botmeshi/app/domain/model"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/infrastructure"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/usecase"
)

func main() {
	logger := log.NewLogger()

	search, err := infrastructure.NewSearchClient(logger)
	if err != nil {
		panic(err.Error())
	}

	notification := infrastructure.NewNotificationClient(logger)
	randomizer := model.NewRandomizer()
	service := service.NewSearchService(search, notification, randomizer, logger)
	uc := usecase.NewSearchUseCase(service, logger)
	lambda.Start(handler.NewSearchHandler(uc, logger).Handler)
}
