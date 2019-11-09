package vo

import "github.com/masakurapa/botmeshi/common/notification/errs"

// Message はテキストメッセージの構造体です
type Message struct {
	value string
}

// NewMessage はテキストメッセージを初期化します
func NewMessage(v string) (Message, error) {
	m := Message{}
	if v == "" {
		return m, errs.NewRequiredArgsError("message")
	}

	m.value = v
	return m, nil
}

// String はテキストメッセージを返します
func (v Message) String() string {
	return v.value
}
