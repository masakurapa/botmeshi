package service

import (
	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/log"
)

// InteractiveService interface
type InteractiveService interface {
	Exec(*api.Parameter) (string, error)
}

type interactiveService struct {
	fnc repository.InvokeFunction
	log log.Logger
}

// NewInteractiveService returns InteractiveService instance
func NewInteractiveService(fnc repository.InvokeFunction, logger log.Logger) InteractiveService {
	return &interactiveService{fnc: fnc, log: logger}
}

// Exec interactive event
func (s *interactiveService) Exec(p *api.Parameter) (msg string, err error) {
	s.log.Start("InteractiveService", "Exec", p)

	switch p.Action.Name {
	case notification.ActionNameCancel:
		msg = "ばいびー"
	case notification.ActionNameGo:
		msg = "`" + p.Action.SelectedOptions[0] + "` に\nごーーーーーーーーーーーーーーーる！！"
	case notification.ActionNameDoNotGo:
		msg = "ざんねん...。"
	case notification.ActionNameSelect:
		err = s.fnc.Exec(&search.Request{
			Target: p.ChannelID,
			Query:  p.Action.SelectedOptions[0],
		})
		if err == nil {
			msg = "`" + p.Action.SelectedOptions[0] + "` でお店を探すよ！\nちょっと時間がかかるからまってくれ！"
		}
	default:
		msg = "キサマ何者だ！"
	}

	s.log.End("InteractiveService", "Exec", msg)
	return msg, err
}
