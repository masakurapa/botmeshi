package infrastructure

import (
	"fmt"

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
	n.log.Start("NotificationClient", "PostMessage", opt)
	n.log.Info("Slack PostMessage parameters", opt)
	err := n.post(opt.Target, slack.MsgOptionText(opt.Message, false))
	n.log.End("NotificationClient", "PostMessage")
	return err
}

func (n *notificationClient) PostRichMessage(opt notification.Option) error {
	n.log.Start("NotificationClient", "PostRichMessage", opt)
	n.log.Info("Slack PostRichMessage parameters", opt)

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

	attachment := slack.Attachment{
		Text:       opt.Message,
		CallbackID: opt.MessageID,
		Color:      opt.Color,
		Actions:    actions,
	}

	n.log.Info("Slack PostRichMessage attachment options", attachment)
	err := n.post(opt.Target, slack.MsgOptionAttachments(attachment))

	n.log.End("NotificationClient", "PostRichMessage")
	return err
}

func (n *notificationClient) post(target string, opt slack.MsgOption) error {
	n.log.Info("Slack PostMessage parameters", opt)
	_, _, err := n.client.PostMessage(target, opt)

	if err != nil {
		n.log.Error("Slack PostMessage error", err)
		return fmt.Errorf("notification error")
	}

	return nil
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
