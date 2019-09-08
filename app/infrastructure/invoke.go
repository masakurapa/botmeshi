package infrastructure

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/masakurapa/botmeshi/app/domain/model/invoke"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/util"
)

type invokeFunction struct {
	client *lambda.Lambda
}

// NewInvokeFunction returns invokeFunction instance
func NewInvokeFunction() repository.InvokeFunction {
	return &invokeFunction{
		client: lambda.New(session.New()),
	}
}

func (f *invokeFunction) Exec(p *invoke.Parameter) error {
	s, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("json encode error")
	}

	_, err = f.client.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(util.InvokeLambdaArn()),
		Payload:        []byte(s),
		InvocationType: aws.String("Event"),
	})

	if err != nil {
		return fmt.Errorf("invoke error")
	}

	return nil
}
