package attachment_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/test/testhelper"
)

func TestSelectBox(t *testing.T) {
	tv1 := testhelper.CreateTextValueVO(t, "text1", "value1")
	tv2 := testhelper.CreateTextValueVO(t, "text2", "value2")
	tvs := testhelper.CreateTextValuesVO(t, tv1, tv2)
	at := attachment.NewSelectBox(tvs)

	t.Run("ActionName().Name()はattachment.NewActionNameSelect().Name()と同じ値を返す", func(t *testing.T) {
		if a := attachment.NewActionNameSelect(); a.Name() != at.ActionName().Name() {
			t.Fatalf("want %q but got %q", a.Name(), at.ActionName().Name())
		}
	})
	t.Run("ActionType().Name()はattachment.NewActionTypeSelectBox().Name()と同じ値を返す", func(t *testing.T) {
		if a := attachment.NewActionTypeSelectBox(); a.Name() != at.ActionType().Name() {
			t.Fatalf("want %q but got %q", a.Name(), at.ActionType().Name())
		}
	})

	opts := at.Options()
	if !t.Run("Options()は初期化パラメータと同じ長さのスライスを返す", func(t *testing.T) {
		tvsopts := tvs.All()
		if len(tvsopts) != len(opts) {
			t.Fatalf("want %d but got %d", len(tvsopts), len(opts))
		}
	}) {
		return
	}
	t.Run("attachment.SelectBoxOption[0].Text()はtv1.Text()と同じ値を返す", func(t *testing.T) {
		if tv1.Text() != opts[0].Text() {
			t.Fatalf("want %q but got %q", tv1.Text(), opts[0].Text())
		}
	})
	t.Run("attachment.SelectBoxOption[0].Value()はtv1.Value()と同じ値を返す", func(t *testing.T) {
		if tv1.Value() != opts[0].Value() {
			t.Fatalf("want %q but got %q", tv1.Value(), opts[0].Value())
		}
	})
	t.Run("attachment.SelectBoxOption[1].Text()はtv2.Text()と同じ値を返す", func(t *testing.T) {
		if tv2.Text() != opts[1].Text() {
			t.Fatalf("want %q but got %q", tv2.Text(), opts[1].Text())
		}
	})
	t.Run("attachment.SelectBoxOption[1].Value()はtv2.Value()と同じ値を返す", func(t *testing.T) {
		if tv2.Value() != opts[1].Value() {
			t.Fatalf("want %q but got %q", tv2.Value(), opts[1].Value())
		}
	})
}
