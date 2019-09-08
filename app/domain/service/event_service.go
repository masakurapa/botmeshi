package service

import (
	"strings"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
)

const (
	ext             = ".json"
	defaultJSON     = "default" + ext
	storageBasePath = ""
	postID          = "menu"
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
}

// NewEventService returns EventService instance
func NewEventService(n repository.Notification) EventService {
	return &eventService{notification: n}
}

// Exec func
func (s *eventService) Exec(p *api.Parameter) error {
	value := s.parse(p.Event.Text)

	if value == "" {
		s.notification.PostMessage(notification.Option{
			Target:  p.ChannelID,
			Message: "探したい駅名を一緒に送って",
		})
		return nil
	}

	return s.interactive(p, value)
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
		MessageID: postID,
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
