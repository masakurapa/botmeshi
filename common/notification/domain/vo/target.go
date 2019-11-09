package vo

import (
	"github.com/masakurapa/botmeshi/common/notification/errs"
)

// Target はメッセージ送信先の構造体です
type Target struct {
	value string
}

// NewTarget はメッセージ送信先の初期化します
func NewTarget(v string) (Target, error) {
	t := Target{}
	if v == "" {
		return t, errs.NewRequiredArgsError("target")
	}

	t.value = v
	return t, nil
}

// String はメッセージ送信先を返します
func (v Target) String() string {
	return v.value
}
