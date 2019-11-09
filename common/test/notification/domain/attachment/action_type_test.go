package attachment_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
)

func TestActionType(t *testing.T) {
	var cases = []struct {
		name   string
		att    attachment.ActionType
		expect string
	}{
		{"ボタンの型のName()はbuttonを返す", attachment.NewActionTypeButton(), "button"},
		{"セレクトボックスの型のName()はbuttonを返す", attachment.NewActionTypeSelectBox(), "select"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if c.expect != c.att.Name() {
				t.Fatalf("want %q but got %q", c.expect, c.att.Name())
			}
		})
	}
}
