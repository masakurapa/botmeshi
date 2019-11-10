package vo

import "github.com/masakurapa/botmeshi/common/faas/errs"

// FunctionName は外部関数名の構造体です
type FunctionName struct {
	value string
}

// NewFunctionName は外部関数名を初期化します
func NewFunctionName(v string) (FunctionName, error) {
	m := FunctionName{}
	if v == "" {
		return m, errs.NewRequiredArgsError("message")
	}

	m.value = v
	return m, nil
}

// String は外部関数名の文字列を返します
func (v FunctionName) String() string {
	return v.value
}
