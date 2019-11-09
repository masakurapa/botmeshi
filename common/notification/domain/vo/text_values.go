package vo

import (
	"github.com/masakurapa/botmeshi/common/notification/errs"
)

// TextValues はText(表示名)とValue(値)のリストの構造体です
type TextValues struct {
	value []TextValue
}

// NewTextValues はText/Valueの構造体を初期化します
func NewTextValues(tv []TextValue) (TextValues, error) {
	tvs := TextValues{}
	if len(tv) == 0 {
		return tvs, errs.NewRequiredArgsError("TextValue values")
	}

	tvs.value = tv
	return tvs, nil
}

// All は全ての要素を返します
func (v TextValues) All() []TextValue {
	return v.value
}
