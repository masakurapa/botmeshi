package vo_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
	"github.com/masakurapa/botmeshi/common/notification/errs"
	"github.com/masakurapa/botmeshi/common/test/testhelper"
)

func TestTextValues_NewTextValues(t *testing.T) {
	tvs := []vo.TextValue{testhelper.CreateTextValueVO(t, "t", "v")}

	t.Run("正常終了時のエラーはnilを返す", func(t *testing.T) {
		if _, err := vo.NewTextValues(tvs); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("All()は引数と同じスライスを返す", func(t *testing.T) {
		if v, _ := vo.NewTextValues(tvs); !reflect.DeepEqual(tvs, v.All()) {
			t.Fatalf("want %+v but got %+v", tvs, v.All())
		}
	})
	t.Run("必須エラー時はerrs.RequiredArgsErrorの構造体を返す", func(t *testing.T) {
		if _, err := vo.NewTextValues([]vo.TextValue{}); errors.Is(err, &errs.RequiredArgsError{}) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
}
