package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/infrastructure"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/usecase"
)

func main() {
	notification := infrastructure.NewNotificationClient()
	service := service.NewEventService(notification)
	uc := usecase.NewEventUseCase(service)
	lambda.Start(handler.NewEventHandler(uc).Handler)
}
