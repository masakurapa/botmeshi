package faas

import "github.com/masakurapa/botmeshi/common/faas/domain/vo"

// Client は外部関数を実行するためのインタフェースです
type Client interface {
	// Invoke は外部関数を実行します
	Invoke(vo.FunctionName, vo.Payload) error
}
