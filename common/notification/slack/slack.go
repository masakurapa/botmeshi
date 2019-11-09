package slack

import (
	"github.com/masakurapa/botmeshi/common/notification"
	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
	"github.com/masakurapa/botmeshi/common/notification/errs"
	slk "github.com/nlopes/slack"
)

// Client はSlack通知用のクライアントのインタフェースです
type Client interface {
	PostMessage(string, ...slk.MsgOption) (string, string, error)
}

// SendSlackClient はSlack通知用のクライアントの構造体です
type SendSlackClient struct {
	Client Client
}

// NewClient は通知用クライアントを初期化します
func NewClient(accessToken string) (notification.Client, error) {
	if accessToken == "" {
		return nil, errs.NewRequiredArgsError("token")
	}

	c := new(SendSlackClient)
	c.Client = slk.New(accessToken)
	return c, nil
}

// PostTextMessage はテキストメッセージ送信を行います
func (c *SendSlackClient) PostTextMessage(channel vo.Target, msg vo.Message) error {
	opt := slk.MsgOptionText(msg.String(), false)
	_, _, err := c.Client.PostMessage(channel.String(), opt)
	return err
}

// PostInteractiveMessage は対話形式のメッセージ送信を行います
func (c *SendSlackClient) PostInteractiveMessage(channel vo.Target, msg vo.Message, att attachment.Attachment) error {
	// セレクトボックス ｰ> ボタンの順番で設定する
	actions := []slk.AttachmentAction{
		c.addSelectAction(att.SelectBox()),
		c.addButtonAction(att.Button()),
	}

	opt := slk.MsgOptionAttachments(slk.Attachment{
		Text:       msg.String(),
		CallbackID: att.ID(),
		Color:      att.Color(),
		Actions:    actions,
	})
	_, _, err := c.Client.PostMessage(channel.String(), opt)
	return err
}

// makeSelectAction はセレクトボックス用のAttachmentを作る
func (c *SendSlackClient) addSelectAction(selectBox attachment.SelectBox) slk.AttachmentAction {
	opts := make([]slk.AttachmentActionOption, len(selectBox.Options()))
	for i, o := range selectBox.Options() {
		opts[i].Text = o.Text()
		opts[i].Value = o.Value()
	}

	return slk.AttachmentAction{
		Name:    selectBox.ActionName().Name(),
		Type:    selectBox.ActionType().Name(),
		Options: opts,
	}
}

// makeButtonAction はボタン用のAttachmentを作る
func (c *SendSlackClient) addButtonAction(btn attachment.Button) slk.AttachmentAction {
	return slk.AttachmentAction{
		Name:  btn.ActionName().Name(),
		Type:  btn.ActionType().Name(),
		Text:  btn.Text(),
		Style: btn.Style(),
	}
}
