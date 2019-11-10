package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lmd "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/masakurapa/botmeshi/common/faas"
	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
)

// Client は別のLambda関数を実行するクライアントのインタフェースです
type Client interface {
	Invoke(*lmd.InvokeInput) (*lmd.InvokeOutput, error)
}

// InvokeClient は別のLambda関数を実行するためのクライアントです
type InvokeClient struct {
	Client Client
}

// NewLamdbaClient はLambda実行用のクライアントを初期化します
func NewLamdbaClient() faas.Client {
	return &InvokeClient{
		Client: lmd.New(session.New()),
	}
}

// Invoke はLambda関数を実行します
func (c *InvokeClient) Invoke(fn vo.FunctionName, p vo.Payload) error {
	params := lmd.InvokeInput{
		FunctionName:   aws.String(fn.String()),
		Payload:        p.Bytes(),
		InvocationType: aws.String("Event"),
	}
	_, err := c.Client.Invoke(&params)
	return err
}
