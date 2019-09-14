package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/infrastructure"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/usecase"
)

func main() {
	search, err := infrastructure.NewSearchClient()
	if err != nil {
		panic(err.Error())
	}

	logger := log.NewLogger()

	notification := infrastructure.NewNotificationClient(logger)
	service := service.NewSearchService(search, notification, logger)
	uc := usecase.NewSearchUseCase(service, logger)
	lambda.Start(handler.NewSearchHandler(uc, logger).Handler)
}
