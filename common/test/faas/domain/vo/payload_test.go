package vo_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
	"github.com/masakurapa/botmeshi/common/faas/errs"
)

func TestPayload(t *testing.T) {
	text := "text"
	t.Run("正常終了時のエラーはnilを返す", func(t *testing.T) {
		if _, err := vo.NewPayload(text); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("String()は引数と同じ文字列を返す", func(t *testing.T) {
		if v, _ := vo.NewPayload(text); text != v.String() {
			t.Fatalf("want %q but got %q", text, v.String())
		}
	})
	t.Run("Bytes()は引数を[]byteにキャストした値を返す", func(t *testing.T) {
		a := []byte(text)
		if v, _ := vo.NewPayload(text); bytes.Compare(a, v.Bytes()) != 0 {
			t.Fatalf("want %v but got %v", a, v.Bytes())
		}
	})
	t.Run("必須エラー時はerrs.RequiredArgsErrorの構造体を返す", func(t *testing.T) {
		if _, err := vo.NewPayload(""); errors.Is(err, &errs.RequiredArgsError{}) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
}
