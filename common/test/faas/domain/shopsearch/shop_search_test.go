package shopserach_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/faas/domain/shopsearch"
	"github.com/masakurapa/botmeshi/common/test/testhelper/faashelper"
)

func TestFunctionName(t *testing.T) {
	callback := faashelper.CreateNotificationCallbackTargetVO(t, "callbacktext")
	query := faashelper.CreateSearchQueryVO(t, "querytext")
	vo := shopsearch.New(callback, query)

	t.Run("Marshal()はcallbackとqueryをJSON文字列に変換して返す", func(t *testing.T) {
		e := "{\"callback\":\"callbacktext\",\"query\":\"querytext\"}"
		if a, _ := vo.Marshal(); e != a {
			t.Fatalf("want %q but got %q", e, a)
		}
	})
}

func TestFunctionName_Unmarshal(t *testing.T) {
	s := "{\"callback\":\"callbacktext\",\"query\":\"querytext\"}"
	p, err := shopsearch.Unmarshal(s)

	if !t.Run("正常終了時はエラーを返却しない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	}) {
		return
	}
	t.Run("Callback().String()はJSONのcallbackの値を返す", func(t *testing.T) {
		e := "callbacktext"
		if a := p.Callback(); e != a.String() {
			t.Fatalf("want %q but got %q", e, a.String())
		}
	})
	t.Run("Query().String()はJSONのqueryの値を返す", func(t *testing.T) {
		e := "querytext"
		if a := p.Query(); e != a.String() {
			t.Fatalf("want %q but got %q", e, a.String())
		}
	})

	var cases = []struct {
		name  string
		param string
	}{
		{"JSON形式ではない文字列の場合はエラーを返す", "text"},
		{"callbackが存在しない場合はエラーを返す", "{\"query\":\"querytext\"}"},
		{"callbackが空文字の場合はエラーを返す", "{\"callback\":\"\",\"query\":\"querytext\"}"},
		{"queryが存在しない場合はエラーを返す", "{\"callback\":\"querytext\"}"},
		{"queryが空文字の場合はエラーを返す", "{\"callback\":\"callbacktext\",\"query\":\"\"}"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if _, err := shopsearch.Unmarshal(c.param); err == nil {
				t.Fatalf("want %t but got %t", true, false)
			}
		})
	}
}
