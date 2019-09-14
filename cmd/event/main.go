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
	logger := log.NewLogger()

	notification := infrastructure.NewNotificationClient(logger)
	service := service.NewEventService(notification, logger)
	uc := usecase.NewEventUseCase(service, logger)
	lambda.Start(handler.NewEventHandler(uc, logger).Handler)
}
