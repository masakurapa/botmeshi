package vo

import "github.com/masakurapa/botmeshi/common/faas/errs"

// NotificationCallbackTarget は外部関数の実行結果通知先の構造体です
type NotificationCallbackTarget struct {
	value string
}

// NewNotificationCallbackTarget は実行結果通知先を初期化します
func NewNotificationCallbackTarget(v string) (NotificationCallbackTarget, error) {
	m := NotificationCallbackTarget{}
	if v == "" {
		return m, errs.NewRequiredArgsError("message")
	}

	m.value = v
	return m, nil
}

// String はテキストメッセージを返します
func (v NotificationCallbackTarget) String() string {
	return v.value
}
