package attachment_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
)

func TestActionName(t *testing.T) {
	var cases = []struct {
		name   string
		att    attachment.ActionName
		expect string
	}{
		{"選択処理の型のName()はactionSelectを返す", attachment.NewActionNameSelect(), "actionSelect"},
		{"選択処理の型のName()はactionSelectを返す", attachment.NewActionNameCancel(), "actionCancel"},
	}

	for _, c := range cases {
		c = c
		t.Run(c.name, func(t *testing.T) {
			if c.expect != c.att.Name() {
				t.Fatalf("want %q but got %q", c.expect, c.att.Name())
			}
		})
	}
}
