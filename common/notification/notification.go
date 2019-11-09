package notification

import (
	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
)

// Notification は各クライアントの初期化を行うためのインタフェースです
type Notification interface {
	// TextMessage はテキストメッセージを送信するためのクライアントを返します
	TextMessage(vo.Target, vo.Message) MessageClient
	// InteractiveSelectMessage はセレクトボックス付きの対話式メッセージを送信するためのクライアントを返します
	InteractiveSelectMessage(vo.Target, vo.Message, vo.TextValues) MessageClient
}

// MessageClient はテキストメッセージを送信するためのインタフェースです
type MessageClient interface {
	// Post はテキストメッセージを送信します
	Post() error
}

// notificationClient は各クライアントを返すための構造体です
type notificationClient struct {
	client Client
}

// New はメッセージ送信クライアントを初期化します
func New(c Client) Notification {
	return &notificationClient{client: c}
}

func (n *notificationClient) TextMessage(target vo.Target, message vo.Message) MessageClient {
	return &textMessageClient{
		client:  n.client,
		target:  target,
		message: message,
	}
}

func (n *notificationClient) InteractiveSelectMessage(target vo.Target, message vo.Message, tv vo.TextValues) MessageClient {
	sb := attachment.NewSelectBox(tv)
	btn := attachment.NewCancelButton()
	att := attachment.NewAttachment(sb, btn)

	return &interactiveMessageClient{
		client:     n.client,
		target:     target,
		message:    message,
		attachment: att,
	}
}

// textMessageClient はテキストメッセージを送信するための構造体です
type textMessageClient struct {
	client  Client
	target  vo.Target
	message vo.Message
}

func (c *textMessageClient) Post() error {
	return c.client.PostTextMessage(c.target, c.message)
}

// interactiveMessageClient は対話式メッセージを送信するための構造体です
type interactiveMessageClient struct {
	client     Client
	target     vo.Target
	message    vo.Message
	attachment attachment.Attachment
}

func (c *interactiveMessageClient) Post() error {
	return c.client.PostInteractiveMessage(c.target, c.message, c.attachment)
}
