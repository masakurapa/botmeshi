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

	fnc := infrastructure.NewInvokeFunction(logger)
	service := service.NewInteractiveService(fnc, logger)
	uc := usecase.NewInteractiveUseCase(service, logger)
	lambda.Start(handler.NewInteractiveHandler(uc, logger).Handler)
}
