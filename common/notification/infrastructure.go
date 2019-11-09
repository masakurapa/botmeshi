package notification

import (
	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
)

// Client は外部にメッセージ送信を行うためのインタフェースです
type Client interface {
	// PostTextMessage はテキストメッセージを送信します
	PostTextMessage(vo.Target, vo.Message) error
	// PostInteractiveMessage は対話式メッセージを送信します
	PostInteractiveMessage(vo.Target, vo.Message, attachment.Attachment) error
}
