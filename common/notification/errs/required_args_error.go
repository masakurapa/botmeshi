package errs

import (
	"fmt"
	"strings"
)

// RequiredArgsError は引数の必須エラーを表す構造体です
type RequiredArgsError struct {
	// 必須となる引数の名称
	args []string
}

// NewRequiredArgsError は引数必須エラーを生成します
func NewRequiredArgsError(args ...string) error {
	return &RequiredArgsError{args: args}
}

func (e *RequiredArgsError) Error() string {
	return fmt.Sprintf("%s is required", strings.Join(e.args, ", "))
}
