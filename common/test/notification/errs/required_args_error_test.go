package errs_test

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/errs"
)

func TestRequiredArgsError(t *testing.T) {
	var cases = []struct {
		name   string
		args   []string
		expect string
	}{
		{"引数が一つの場合に期待するメッセージを返す", []string{"arg1"}, "arg1 is required"},
		{"引数が複数の場合に期待するメッセージを返す", []string{"arg1", "arg2", "arg3"}, "arg1, arg2, arg3 is required"},
	}

	for _, c := range cases {
		c := c
		err := errs.NewRequiredArgsError(c.args...)
		t.Run(c.name, func(t *testing.T) {
			a := err.Error()
			if c.expect != a {
				t.Fatalf("want %q but got %q", c.expect, a)
			}
		})
	}
}
