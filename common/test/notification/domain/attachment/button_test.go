package attachment_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
)

func TestButton_CancelButton(t *testing.T) {
	at := attachment.NewCancelButton()

	t.Run("Text()はキャンセルを返す", func(t *testing.T) {
		e := "キャンセル"
		if e != at.Text() {
			t.Fatalf("want %q but got %q", e, at.Text())
		}
	})
	t.Run("Style()はdangerを返す", func(t *testing.T) {
		e := "danger"
		if e != at.Style() {
			t.Fatalf("want %q but got %q", e, at.Style())
		}
	})
	t.Run("ActionName().Name()はattachment.NewActionNameCancel().Name()と同じ値を返す", func(t *testing.T) {
		if a := attachment.NewActionNameCancel(); a.Name() != at.ActionName().Name() {
			t.Fatalf("want %q but got %q", a.Name(), at.ActionName().Name())
		}
	})
	t.Run("ActionType().Name()はattachment.NewActionTypeButton().Name()と同じ値を返す", func(t *testing.T) {
		if a := attachment.NewActionTypeButton(); a.Name() != at.ActionType().Name() {
			t.Fatalf("want %q but got %q", a.Name(), at.ActionType().Name())
		}
	})
}
