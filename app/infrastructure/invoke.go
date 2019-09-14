package infrastructure

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/util"
)

type invokeFunction struct {
	client *lambda.Lambda
	log    log.Logger
}

// NewInvokeFunction returns invokeFunction instance
func NewInvokeFunction(logger log.Logger) repository.InvokeFunction {
	return &invokeFunction{
		client: lambda.New(session.New()),
		log:    logger,
	}
}

func (f *invokeFunction) Exec(p *search.Request) error {
	f.log.Start("InvokeFunction", "Exec", p)

	s, err := json.Marshal(p)
	if err != nil {
		f.log.Error("JSON parse error", err)
		return err
	}

	params := lambda.InvokeInput{
		FunctionName:   aws.String(util.InvokeLambdaArn()),
		Payload:        []byte(s),
		InvocationType: aws.String("Event"),
	}

	f.log.Info("Lambda invoke parameters", params)
	_, err = f.client.Invoke(&params)

	if err != nil {
		f.log.Error("Lambda invoke error", err)
		return err
	}

	f.log.End("InvokeFunction", "Exec")
	return nil
}
