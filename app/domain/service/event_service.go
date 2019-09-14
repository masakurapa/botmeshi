package service

import (
	"strings"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
)

const (
	ext                = ".json"
	defaultJSON        = "default" + ext
	storageBasePath    = ""
	postEventMessageID = "menu"
)

var menus = []string{
	"ラーメン", "肉", "魚", "定食", "カレー", "和食", "中華",
}

// EventService interface
type EventService interface {
	Exec(*api.Parameter) error
}

type eventService struct {
	notification repository.Notification
	log          log.Logger
}

// NewEventService returns EventService instance
func NewEventService(n repository.Notification, logger log.Logger) EventService {
	return &eventService{notification: n, log: logger}
}

// Exec func
func (s *eventService) Exec(p *api.Parameter) error {
	s.log.Start("EventService", "Exec", p)

	value := s.parse(p.Event.Text)
	if value == "" {
		s.log.Error("Event text is required")

		s.notification.PostMessage(notification.Option{
			Target:  p.ChannelID,
			Message: "探したい駅名を一緒に送って",
		})
		return nil
	}

	err := s.interactive(p, value)
	s.log.End("EventService", "Exec")
	return err
}

func (s *eventService) parse(text string) string {
	// 先頭12文字はメンション用の固定文字なのでいらない
	if len(text) < 12 {
		return ""
	}
	return strings.TrimSpace(text[12:])
}

func (s *eventService) interactive(p *api.Parameter, text string) error {
	var selectOpt []notification.SelectOption
	for _, g := range menus {
		selectOpt = append(selectOpt, notification.SelectOption{
			Text:  g,
			Value: text + " " + g,
		})
	}

	return s.notification.PostRichMessage(notification.Option{
		Target:    p.ChannelID,
		Message:   text + " で何が食べたい？",
		MessageID: postEventMessageID,
		Color:     "#ff6633",
		RichMessageOptions: []notification.RichMessageOption{
			{
				ActionName:    notification.ActionNameSelect,
				ActionType:    notification.ActionTypeSelect,
				SelectOptions: selectOpt,
			},
			{
				ActionName: notification.ActionNameCancel,
				ActionType: notification.ActionTypeButton,
				Text:       "やめる",
				Style:      "danger",
			},
		},
	})
}
