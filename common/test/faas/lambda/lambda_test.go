package lambda_test

import (
	"bytes"
	"fmt"
	"testing"

	lmd "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/masakurapa/botmeshi/common/faas"
	"github.com/masakurapa/botmeshi/common/faas/lambda"
	"github.com/masakurapa/botmeshi/common/test/testhelper/faashelper"
)

type mockLambdaInvokeClient struct {
	faas.Client
	mockInvoke func(*lmd.InvokeInput) (*lmd.InvokeOutput, error)
}

func (m *mockLambdaInvokeClient) Invoke(input *lmd.InvokeInput) (*lmd.InvokeOutput, error) {
	return m.mockInvoke(input)
}

func TestNewLamdbaClient(t *testing.T) {
	c := lambda.NewLamdbaClient()
	t.Run("faas.Clientの実装クラスを返す", func(t *testing.T) {
		if _, ok := c.(faas.Client); !ok {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
}

func TestInvoke(t *testing.T) {
	fn := faashelper.CreateFunctionNameVO(t, "lambdafunction")
	pld := faashelper.CreatePayloadVO(t, "payload")

	var actInput *lmd.InvokeInput

	c := lambda.InvokeClient{
		Client: &mockLambdaInvokeClient{
			mockInvoke: func(i *lmd.InvokeInput) (*lmd.InvokeOutput, error) {
				actInput = i
				return nil, nil
			},
		},
	}
	err := c.Invoke(fn, pld)

	t.Run("Invoke()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("Invoke()の引数のFunctionNameはFunctionName.String()と等しい", func(t *testing.T) {
		if fn.String() != *actInput.FunctionName {
			t.Fatalf("want %q but got %q", fn.String(), *actInput.FunctionName)
		}
	})
	t.Run("Invoke()の引数のPayloadはPayload.Bytes()と等しい", func(t *testing.T) {
		if bytes.Compare(pld.Bytes(), actInput.Payload) != 0 {
			t.Fatalf("want %v but got %v", pld.Bytes(), actInput.Payload)
		}
	})
	t.Run("Invoke()の引数のInvocationTypeはEventと等しい", func(t *testing.T) {
		e := "Event"
		if e != *actInput.InvocationType {
			t.Fatalf("want %q but got %q", e, *actInput.InvocationType)
		}
	})

	t.Run("Post()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c := lambda.InvokeClient{
			Client: &mockLambdaInvokeClient{
				mockInvoke: func(i *lmd.InvokeInput) (*lmd.InvokeOutput, error) {
					return nil, mockError
				},
			},
		}
		if err := c.Invoke(fn, pld); err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}
