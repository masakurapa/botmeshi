package faas_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/common/faas"
	"github.com/masakurapa/botmeshi/common/faas/domain/shopsearch"
	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
	"github.com/masakurapa/botmeshi/common/test/testhelper/faashelper"
)

type mockClient struct {
	faas.Client
	mockInvoke func(vo.FunctionName, vo.Payload) error
}

func (m *mockClient) Invoke(fn vo.FunctionName, p vo.Payload) error {
	return m.mockInvoke(fn, p)
}

func TestFaas_InvokeShopSearch(t *testing.T) {
	callback := faashelper.CreateNotificationCallbackTargetVO(t, "callback")
	query := faashelper.CreateSearchQueryVO(t, "query text")
	sp := shopsearch.New(callback, query)

	fncName := "func name"

	var actFn vo.FunctionName
	var actPld vo.Payload

	c, err := faas.New(&mockClient{
		mockInvoke: func(fn vo.FunctionName, p vo.Payload) error {
			actFn = fn
			actPld = p
			return nil
		},
	}, fncName)

	if !t.Run("New()は関数名が指定されている場合はエラーを返さない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	}) {
		return
	}

	err = c.InvokeShopSearch(sp)
	t.Run("Invoke()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("Invoke()のvo.FunctionNameはNew()の引数と同じ", func(t *testing.T) {
		if fncName != actFn.String() {
			t.Fatalf("want %q but got %q", fncName, actFn.String())
		}
	})
	t.Run("Invoke()のvo.Payloadはshopsearch.Marchal()と同じ", func(t *testing.T) {
		b, _ := sp.Marshal()
		pld := faashelper.CreatePayloadVO(t, b)
		if bytes.Compare(pld.Bytes(), actPld.Bytes()) != 0 {
			t.Fatalf("want %v but got %v", pld.Bytes(), actPld.Bytes())
		}
	})

	t.Run("New()は関数名が指定されていない場合にエラー返す", func(t *testing.T) {
		if _, err := faas.New(&mockClient{}, ""); err == nil {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
	t.Run("Invoke()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c, _ := faas.New(&mockClient{
			mockInvoke: func(fn vo.FunctionName, p vo.Payload) error {
				return mockError
			},
		}, fncName)
		err := c.InvokeShopSearch(sp)
		if err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}
