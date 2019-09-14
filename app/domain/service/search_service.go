package service

import (
	"math/rand"

	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/util"
)

const (
	radius              uint = 500
	radiusBuffer        uint = 100
	shopMax                  = 5
	postSearchMessageID      = "shop"
)

var randInt = rand.Intn

// SearchService interface
type SearchService interface {
	SearchStation(*search.Query) *search.Station
	SearchShops(*search.Query, *search.Station) []search.Shop
	NoticeSuccess(*search.Request, []search.Shop) error
	NoticeError(*search.Request, string)
}

type searchService struct {
	client       repository.Search
	notification repository.Notification
}

// NewSearchService returns SearchService instance
func NewSearchService(s repository.Search, n repository.Notification) SearchService {
	return &searchService{
		client:       s,
		notification: n,
	}
}

// SearchStation func
func (s *searchService) SearchStation(q *search.Query) *search.Station {
	return s.client.Station(q)
}

// SearchShops func
func (s *searchService) SearchShops(q *search.Query, station *search.Station) []search.Shop {
	var p search.SearchShopsQuery
	if station == nil {
		// 駅がないときはエリア名 + ジャンル名
		p.Query = q.AreaName + " " + q.Genre
	} else {
		p.Query = q.Genre
		p.Radius = radius
		p.Location = &station.Location
	}

	shops := s.client.Shops(&p)
	if len(shops) == 0 {
		return shops
	}

	return s.getShopURL(q, s.random(s.filterNearShops(station, shops)))
}

// NoticeSuccess func
func (s *searchService) NoticeSuccess(r *search.Request, shops []search.Shop) error {
	var opts []notification.SelectOption
	text := "お店を見つけたよ！！\n```\n"

	for _, shop := range shops {
		text += shop.Name + "\n"
		if shop.URL != "" {
			text += shop.URL + "\n"
		}

		opts = append(opts, notification.SelectOption{
			Text:  shop.Name,
			Value: shop.Name,
		})
	}

	// 店の情報だけまず送る
	err := s.notification.PostMessage(notification.Option{
		Target:  r.Target,
		Message: text + "```\n",
	})
	if err != nil {
		return err
	}

	// interactive
	opt := notification.Option{
		Target:    r.Target,
		Message:   "いいお店は見つかったかな？",
		MessageID: postSearchMessageID,
		Color:     "#ff6633",
		RichMessageOptions: []notification.RichMessageOption{
			{
				ActionName:    notification.ActionNameGo,
				ActionType:    notification.ActionTypeSelect,
				SelectOptions: opts,
			},
			{
				ActionName: notification.ActionNameDoNotGo,
				ActionType: notification.ActionTypeButton,
				Text:       "いい店はなかった",
				Style:      notification.ButtonStyleDanger,
			},
		},
	}

	// 店の情報は送信できているのでこっちの通知は失敗してもエラーにしない
	s.notification.PostRichMessage(opt)
	return nil
}

// NoticeError func
func (s *searchService) NoticeError(r *search.Request, msg string) {
	s.notification.PostMessage(notification.Option{
		Target:  r.Target,
		Message: msg,
	})
}

// filterNearShops 駅近くの店だけにフィルタリング
func (s *searchService) filterNearShops(station *search.Station, shops []search.Shop) []search.Shop {
	// 駅情報がないと何もできない
	if station == nil {
		return shops
	}

	// 場所検索時の半径に少しバッファをもたせる
	r := float64(radius + radiusBuffer)
	var ret []search.Shop
	for _, p := range shops {
		m := util.Distance(station.Lat, station.Lng, p.Location.Lat, p.Location.Lng)
		if m > r {
			continue
		}
		ret = append(ret, p)
	}
	return ret
}

// random 店をランダムで返す
func (s *searchService) random(shops []search.Shop) []search.Shop {
	if len(shops) <= shopMax {
		return shops
	}

	n := len(shops)
	for i := n - 1; i >= 0; i-- {
		j := randInt(i + 1)
		shops[i], shops[j] = shops[j], shops[i]
	}
	return shops[:shopMax]
}

// getShopURL 店のURLを取得する
func (s *searchService) getShopURL(q *search.Query, shops []search.Shop) []search.Shop {
	for i, shop := range shops {
		ret := s.client.Shop(&search.SearchShopsQuery{
			Query: q.AreaName + " " + q.Genre + " " + shop.Name,
		})
		if ret == nil {
			continue
		}
		shops[i].URL = ret.URL
	}
	return shops
}
