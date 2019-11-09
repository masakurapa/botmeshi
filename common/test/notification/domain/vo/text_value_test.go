package vo_test

import (
	"errors"
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
	"github.com/masakurapa/botmeshi/common/notification/errs"
)

func TestTextValue(t *testing.T) {
	text := "text"
	value := "value"

	t.Run("正常終了時のエラーはnilを返す", func(t *testing.T) {
		if _, err := vo.NewTextValue(text, value); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("Text()は第一引数と同じ文字列を返す", func(t *testing.T) {
		if v, _ := vo.NewTextValue(text, value); text != v.Text() {
			t.Fatalf("want %q but got %q", text, v.Text())
		}
	})
	t.Run("Value()は第二引数と同じ文字列を返す", func(t *testing.T) {
		if v, _ := vo.NewTextValue(text, value); value != v.Value() {
			t.Fatalf("want %q but got %q", value, v.Value())
		}
	})

	var cases = []struct {
		name  string
		text  string
		value string
	}{
		{"textの必須エラー時はerrs.RequiredArgsErrorの構造体を返す", "", value},
		{"valueの必須エラー時はerrs.RequiredArgsErrorの構造体を返す", text, ""},
		{"text/valueの必須エラー時はerrs.RequiredArgsErrorの構造体を返す", "", ""},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewTextValue(c.text, c.value); errors.Is(err, &errs.RequiredArgsError{}) {
				t.Fatalf("want %t but got %t", true, false)
			}
		})
	}
}
