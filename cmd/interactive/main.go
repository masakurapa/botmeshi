package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/infrastructure"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/usecase"
)

func main() {
	fnc := infrastructure.NewInvokeFunction()
	service := service.NewInteractiveService(fnc)
	uc := usecase.NewInteractiveUseCase(service)
	lambda.Start(handler.NewInteractiveHandler(uc).Handler)
}
