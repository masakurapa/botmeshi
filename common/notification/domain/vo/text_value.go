package vo

import "github.com/masakurapa/botmeshi/common/notification/errs"

// TextValue はText(表示名)とValue(値)の構造体です
type TextValue struct {
	text  string
	value string
}

// NewTextValue はText/Valueの構造体を初期化します
func NewTextValue(t, v string) (TextValue, error) {
	tv := TextValue{}
	if t == "" || v == "" {
		return tv, errs.NewRequiredArgsError("text", "value")
	}

	tv.text = t
	tv.value = v
	return tv, nil
}

// Text は表示名の文字列を返します
func (v TextValue) Text() string {
	return v.text
}

// Value は値の文字列を返します
func (v TextValue) Value() string {
	return v.value
}
