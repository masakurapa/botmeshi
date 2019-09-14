package usecase

import (
	"fmt"
	"strings"

	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/log"
)

const (
	// クエリの区切り文字
	querySep = " "
)

// SearchUseCase interface
type SearchUseCase interface {
	Validate(*search.Request) error
	Exec(*search.Request) error
}

type searchUseCase struct {
	service service.SearchService
	log     log.Logger
}

// NewSearchUseCase return SearchUseCase instance
func NewSearchUseCase(s service.SearchService, logger log.Logger) SearchUseCase {
	return &searchUseCase{service: s, log: logger}
}

func (uc *searchUseCase) Validate(p *search.Request) error {
	if strings.TrimSpace(p.Target) == "" {
		return fmt.Errorf("target is required")
	}
	if strings.TrimSpace(p.Query) == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

func (uc *searchUseCase) Exec(p *search.Request) error {
	query := uc.parse(p.Query)
	if query == nil {
		return fmt.Errorf("invalid query: " + p.Query)
	}

	station := uc.service.SearchStation(query)

	shops := uc.service.SearchShops(query, station)
	if len(shops) == 0 {
		uc.service.NoticeError(p, "お店が見つからなかったよ")
		return fmt.Errorf("shop not found")
	}

	if err := uc.service.NoticeSuccess(p, shops); err != nil {
		uc.service.NoticeError(p, "お店の通知ができなかったよ")
		return fmt.Errorf("notice error")
	}

	return nil
}

func (uc *searchUseCase) parse(query string) *search.Query {
	s := strings.Split(query, querySep)
	if len(s) == 1 {
		return nil
	}

	i := len(s) - 1
	return &search.Query{
		AreaName: strings.Join(s[0:i], querySep),
		Genre:    s[i],
	}
}
