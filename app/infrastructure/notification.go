package infrastructure

import (
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
	"github.com/masakurapa/botmeshi/app/util"
	"github.com/nlopes/slack"
)

type notificationClient struct {
	client *slack.Client
	log    log.Logger
}

// NewNotificationClient returns Notification instance
func NewNotificationClient(logger log.Logger) repository.Notification {
	return &notificationClient{
		client: slack.New(util.BotAccessToken()),
		log:    logger,
	}
}

func (n *notificationClient) PostMessage(opt notification.Option) error {
	return n.post(opt.Target, slack.MsgOptionText(opt.Message, false))
}

func (n *notificationClient) PostRichMessage(opt notification.Option) error {
	var actions []slack.AttachmentAction
	for _, o := range opt.RichMessageOptions {
		var action slack.AttachmentAction

		switch o.ActionType {
		case notification.ActionTypeButton:
			action = slack.AttachmentAction{
				Name:  o.ActionName,
				Type:  o.ActionType,
				Text:  o.Text,
				Style: o.Style,
			}
		case notification.ActionTypeSelect:
			action = slack.AttachmentAction{
				Name:    o.ActionName,
				Type:    o.ActionType,
				Options: n.makeAttachmentActionOption(o.SelectOptions),
			}
		}

		actions = append(actions, action)
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
