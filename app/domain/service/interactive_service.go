package service

import (
	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
)

// InteractiveService interface
type InteractiveService interface {
	Exec(*api.Parameter) (string, error)
}

type interactiveService struct {
	fnc repository.InvokeFunction
}

// NewInteractiveService returns InteractiveService instance
func NewInteractiveService(fnc repository.InvokeFunction) InteractiveService {
	return &interactiveService{
		fnc: fnc,
	}
}

// Exec interactive event
func (s *interactiveService) Exec(p *api.Parameter) (msg string, err error) {
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

	return msg, err
}

func (s *interactiveService) invoke() {

}
