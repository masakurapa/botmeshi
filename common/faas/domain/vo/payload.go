package vo

import "github.com/masakurapa/botmeshi/common/faas/errs"

// Payload は外部関数に渡すデータ本体の構造体です
type Payload struct {
	value string
}

// NewPayload はPayloadを初期化します
func NewPayload(v string) (Payload, error) {
	m := Payload{}
	if v == "" {
		return m, errs.NewRequiredArgsError("message")
	}

	m.value = v
	return m, nil
}

// String はPayloadの文字列を返します
func (v Payload) String() string {
	return v.value
}

// Bytes はPayloadのバイトのスライスを返します
func (v Payload) Bytes() []byte {
	return []byte(v.value)
}
