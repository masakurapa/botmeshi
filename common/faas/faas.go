package faas

import (
	"github.com/masakurapa/botmeshi/common/faas/domain/shopsearch"
	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
)

// Faas は各クライアントの初期化を行うためのインタフェースです
type Faas interface {
	// InvokeShopSearch は店検索の外部関数を実行します
	InvokeShopSearch(shopsearch.Params) error
}

// faasClient は各クライアントを返すための構造体です
type faasClient struct {
	client Client
	fn     vo.FunctionName
}

// New はメッセージ送信クライアントを初期化します
func New(c Client, fn string) (Faas, error) {
	vfn, err := vo.NewFunctionName(fn)
	if err != nil {
		return nil, err
	}
	return &faasClient{client: c, fn: vfn}, nil
}

func (c *faasClient) InvokeShopSearch(p shopsearch.Params) error {
	b, err := p.Marshal()
	if err != nil {
		return err
	}
	pld, err := vo.NewPayload(b)
	if err != nil {
		return err
	}

	return c.client.Invoke(c.fn, pld)
}
