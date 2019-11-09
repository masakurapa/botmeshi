package attachment_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/test/testhelper"
)

func TestAttachment(t *testing.T) {
	tvs := testhelper.CreateTextValuesVO(t, testhelper.CreateTextValueVO(t, "t", "v"))
	sb := attachment.NewSelectBox(tvs)
	btn := attachment.NewCancelButton()
	att := attachment.NewAttachment(sb, btn)

	t.Run("Color()は#ff6633を返す", func(t *testing.T) {
		e := "#ff6633"
		if e != att.Color() {
			t.Fatalf("want %q but got %q", e, att.Color())
		}
	})
	t.Run("SelectBox()は引数のattachment.SelectBoxの構造体を返す", func(t *testing.T) {
		if !reflect.DeepEqual(sb, att.SelectBox()) {
			t.Fatalf("want %v but got %v", sb, att.SelectBox())
		}
	})
	t.Run("Button()は引数のattachment.Buttonの構造体を返す", func(t *testing.T) {
		if !reflect.DeepEqual(btn, att.Button()) {
			t.Fatalf("want %v but got %v", sb, att.Button())
		}
	})
	t.Run("ID()はprefix + ActionName + ActionTypeの文字列を返す", func(t *testing.T) {
		e := "notice" + strings.Title(sb.ActionName().Name()) + strings.Title(sb.ActionType().Name())
		if e != att.ID() {
			t.Fatalf("want %q but got %q", e, att.ID())
		}
	})
}
