package infrastructure

import (
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/util"
	"github.com/nlopes/slack"
)

type notificationClient struct {
	client *slack.Client
}

// NewNotificationClient returns Notification instance
func NewNotificationClient() repository.Notification {
	return &notificationClient{
		client: slack.New(util.BotAccessToken()),
	}
}

func (n *notificationClient) PostMessage(opt notification.Option) error {
	return n.post(opt.Target, slack.MsgOptionText(opt.Message, false))
}

func (n *notificationClient) PostRichMessage(opt notification.Option) error {
	var actions []slack.AttachmentAction
	for _, o := range opt.RichMessageOptions {
		action := slack.AttachmentAction{
			Name: o.ActionName,
			Type: o.ActionType,
		}

		switch o.ActionType {
		case notification.ActionTypeButton:
			action.Text = o.Text
			action.Style = o.Style
		case notification.ActionTypeSelect:
			action.Options = n.makeAttachmentActionOption(o.SelectOptions)
		}

		actions = append(actions, slack.AttachmentAction{
			Name: o.ActionType,
		})
	}

	return n.post(opt.Target, slack.MsgOptionAttachments(slack.Attachment{
		Text:       opt.Message,
		CallbackID: opt.MessageID,
		Color:      opt.Color,
		Actions:    actions,
	}))
}

func (n *notificationClient) post(target string, opt slack.MsgOption) error {
	_, _, err := n.client.PostMessage(target, opt)
	return err
}

func (n *notificationClient) makeAttachmentActionOption(options []notification.SelectOption) []slack.AttachmentActionOption {
	var opts []slack.AttachmentActionOption
	for _, o := range options {
		opts = append(opts, slack.AttachmentActionOption{
			Text:  o.Text,
			Value: o.Value,
		})
	}
	return opts
}
