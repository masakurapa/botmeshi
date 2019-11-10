package vo_test

import (
	"errors"
	"testing"

	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
	"github.com/masakurapa/botmeshi/common/faas/errs"
)

func TestNotificationCallbackTarget(t *testing.T) {
	text := "text"
	t.Run("正常終了時のエラーはnilを返す", func(t *testing.T) {
		if _, err := vo.NewNotificationCallbackTarget(text); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("String()は引数と同じ文字列を返す", func(t *testing.T) {
		if v, _ := vo.NewNotificationCallbackTarget(text); text != v.String() {
			t.Fatalf("want %q but got %q", text, v.String())
		}
	})
	t.Run("必須エラー時はerrs.RequiredArgsErrorの構造体を返す", func(t *testing.T) {
		if _, err := vo.NewNotificationCallbackTarget(""); errors.Is(err, &errs.RequiredArgsError{}) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
}
