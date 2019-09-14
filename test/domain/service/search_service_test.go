package service

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/stretchr/testify/assert"
)

type testSearchMock struct {
	repository.Search
	stationMock func(*search.Query) *search.Station
	shopMock    func(*search.SearchShopsQuery) *search.Shop
	shopsMock   func(*search.SearchShopsQuery) []search.Shop
}

func (t *testSearchMock) Station(q *search.Query) *search.Station {
	return t.stationMock(q)
}
func (t *testSearchMock) Shop(q *search.SearchShopsQuery) *search.Shop {
	return t.shopMock(q)
}
func (t *testSearchMock) Shops(q *search.SearchShopsQuery) []search.Shop {
	return t.shopsMock(q)
}

type testRandomizerMock struct {
	repository.Randomizer
	intnMock func() int
}

func (t *testRandomizerMock) Intn(int) int {
	return t.intnMock()
}

type testSearchNotificationMock struct {
	repository.Notification
	postMessageMock     func(option notification.Option) error
	postRichMessageMock func(option notification.Option) error
}

func (t *testSearchNotificationMock) PostMessage(o notification.Option) error {
	return t.postMessageMock(o)
}
func (t *testSearchNotificationMock) PostRichMessage(o notification.Option) error {
	return t.postRichMessageMock(o)
}

func TestNewSearchService(t *testing.T) {
	func() {
		s := service.NewSearchService(&testSearchMock{}, &testSearchNotificationMock{}, &testRandomizerMock{}, &loggerMock{})
		_, ok := s.(service.SearchService)
		assert.True(t, ok)
	}()
}

func TestSearchService_SearchStation(t *testing.T) {
	s := testSearchMock{}
	n := &testSearchNotificationMock{}

	func(s testSearchMock) {
		expect := &search.Station{}
		s.stationMock = func(q *search.Query) *search.Station {
			assert.Equal(t, "東京", q.AreaName)
			assert.Equal(t, "ラーメン", q.Genre)
			return expect
		}
		actual := service.NewSearchService(&s, n, &testRandomizerMock{}, &loggerMock{}).SearchStation(&search.Query{
			AreaName: "東京",
			Genre:    "ラーメン",
		})
		assert.Equal(t, expect, actual)
	}(s)

	// データ0件
	func(s testSearchMock) {
		s.stationMock = func(*search.Query) *search.Station { return nil }

		actual := service.NewSearchService(&s, n, &testRandomizerMock{}, &loggerMock{}).SearchStation(&search.Query{})
		assert.Nil(t, actual)
	}(s)
}

func TestSearchService_SearchShops(t *testing.T) {
	s := testSearchMock{}
	n := &testSearchNotificationMock{}

	station := &search.Station{Name: "東京駅", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}}
	shops := []search.Shop{
		{Name: "東京の店1", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "東京の店2", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "東京の店3", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "東京の店4", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "東京の店5", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "東京の店6", Location: search.Location{Lat: 35.6812362, Lng: 139.7649361}},
		{Name: "渋谷の店1", Location: search.Location{Lat: 35.6580339, Lng: 139.6994471}},
		{Name: "渋谷の店2", Location: search.Location{Lat: 35.6580339, Lng: 139.6994471}},
	}

	shopMockValues := []search.Shop{
		{URL: "http://localhost/shop1"},
		{URL: "http://localhost/shop2"},
		{URL: "http://localhost/shop3"},
		{URL: "http://localhost/shop4"},
		{URL: ""},
	}

	// 正常系（駅情報あり
	func(s testSearchMock) {
		s.shopsMock = func(q *search.SearchShopsQuery) []search.Shop {
			assert.Equal(t, "ラーメン", q.Query)
			assert.Equal(t, uint(500), q.Radius)
			assert.Equal(t, &station.Location, q.Location)
			return shops
		}

		cnt := 0
		s.shopMock = func(q *search.SearchShopsQuery) *search.Shop {
			assert.Equal(t, "東京 ラーメン "+shops[cnt].Name, q.Query)

			s := shopMockValues[cnt]
			cnt++
			if s.URL == "" {
				return nil
			}
			return &s
		}

		num := 0
		rand := &testRandomizerMock{
			intnMock: func() int {
				n := num
				num++
				return n
			},
		}

		actual := service.NewSearchService(&s, n, rand, &loggerMock{}).SearchShops(&search.Query{
			AreaName: "東京",
			Genre:    "ラーメン",
		}, station)

		if assert.Equal(t, 5, len(actual)) {
			assert.Equal(t, search.Shop{
				Name:     "東京の店1",
				URL:      "http://localhost/shop1",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[0])
			assert.Equal(t, search.Shop{
				Name:     "東京の店2",
				URL:      "http://localhost/shop2",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[1])
			assert.Equal(t, search.Shop{
				Name:     "東京の店3",
				URL:      "http://localhost/shop3",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[2])
			assert.Equal(t, search.Shop{
				Name:     "東京の店4",
				URL:      "http://localhost/shop4",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[3])
			assert.Equal(t, search.Shop{
				Name:     "東京の店5",
				URL:      "",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[4])
		}
	}(s)

	// 正常系（駅情報なし
	func(s testSearchMock) {
		s.shopsMock = func(q *search.SearchShopsQuery) []search.Shop {
			assert.Equal(t, "東京 ラーメン", q.Query)
			assert.Nil(t, q.Location)
			return shops
		}

		cnt := 0
		s.shopMock = func(q *search.SearchShopsQuery) *search.Shop {
			assert.Equal(t, "東京 ラーメン "+shops[cnt].Name, q.Query)

			s := shopMockValues[cnt]
			cnt++
			if s.URL == "" {
				return nil
			}
			return &s
		}

		num := 0
		rand := &testRandomizerMock{
			intnMock: func() int {
				n := num
				num++
				return n
			},
		}

		actual := service.NewSearchService(&s, n, rand, &loggerMock{}).SearchShops(&search.Query{
			AreaName: "東京",
			Genre:    "ラーメン",
		}, nil)

		if assert.Equal(t, 5, len(actual)) {
			assert.Equal(t, search.Shop{
				Name:     "東京の店1",
				URL:      "http://localhost/shop1",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[0])
			assert.Equal(t, search.Shop{
				Name:     "東京の店2",
				URL:      "http://localhost/shop2",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[1])
			assert.Equal(t, search.Shop{
				Name:     "東京の店3",
				URL:      "http://localhost/shop3",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[2])
			assert.Equal(t, search.Shop{
				Name:     "東京の店4",
				URL:      "http://localhost/shop4",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[3])
			assert.Equal(t, search.Shop{
				Name:     "東京の店5",
				URL:      "",
				Location: search.Location{Lat: 35.6812362, Lng: 139.7649361},
			}, actual[4])
		}
	}(s)

	// 異常系（店無し
	func(s testSearchMock) {
		expect := []search.Shop{}
		s.shopsMock = func(*search.SearchShopsQuery) []search.Shop { return expect }
		actual := service.NewSearchService(&s, n, &testRandomizerMock{}, &loggerMock{}).SearchShops(&search.Query{}, nil)
		assert.Equal(t, expect, actual)
	}(s)
}

func TestSearchService_NoticeSuccess(t *testing.T) {
	s := &testSearchMock{}
	n := testSearchNotificationMock{}

	shops := []search.Shop{
		{Name: "shop1", URL: "http://localhost/shop1"},
		{Name: "shop2", URL: ""},
	}

	// 正常系
	func(n testSearchNotificationMock) {
		n.postMessageMock = func(o notification.Option) error {
			msg := "お店を見つけたよ！！\n```\n"
			msg += "shop1\nhttp://localhost/shop1\n"
			msg += "shop2\n"
			msg += "```\n"

			assert.Equal(t, "12345", o.Target)
			assert.Equal(t, msg, o.Message)
			return nil
		}
		n.postRichMessageMock = func(opt notification.Option) error {
			assert.Equal(t, "12345", opt.Target)
			assert.Equal(t, "いいお店は見つかったかな？", opt.Message)
			assert.Equal(t, "shop", opt.MessageID)
			assert.Equal(t, "#ff6633", opt.Color)
			assert.Equal(t, []notification.RichMessageOption{
				{
					ActionName: notification.ActionNameGo,
					ActionType: notification.ActionTypeSelect,
					SelectOptions: []notification.SelectOption{
						{Text: "shop1", Value: "shop1"},
						{Text: "shop2", Value: "shop2"},
					},
				},
				{
					ActionName: notification.ActionNameDoNotGo,
					ActionType: notification.ActionTypeButton,
					Text:       "いい店はなかった",
					Style:      notification.ButtonStyleDanger,
				},
			}, opt.RichMessageOptions)
			return nil
		}

		err := service.NewSearchService(s, &n, &testRandomizerMock{}, &loggerMock{}).NoticeSuccess(&search.Request{
			Target: "12345",
		}, shops)
		assert.Nil(t, err)
	}(n)

	// 店の情報通知に失敗
	func(n testSearchNotificationMock) {
		n.postMessageMock = func(o notification.Option) error { return fmt.Errorf("post message error") }

		err := service.NewSearchService(s, &n, &testRandomizerMock{}, &loggerMock{}).NoticeSuccess(&search.Request{
			Target: "12345",
		}, shops)
		assert.NotNil(t, err)
		assert.Equal(t, "post message error", err.Error())
	}(n)

	// interactiveメッセージの通知に失敗
	func(n testSearchNotificationMock) {
		n.postMessageMock = func(o notification.Option) error { return nil }
		n.postRichMessageMock = func(o notification.Option) error { return fmt.Errorf("post rich message error") }

		err := service.NewSearchService(s, &n, &testRandomizerMock{}, &loggerMock{}).NoticeSuccess(&search.Request{
			Target: "12345",
		}, shops)
		assert.Nil(t, err)
	}(n)
}

func TestSearchService_NoticeError(t *testing.T) {
	s := &testSearchMock{}
	n := testSearchNotificationMock{}

	func() {
		n.postMessageMock = func(o notification.Option) error {
			assert.Equal(t, "12345", o.Target)
			assert.Equal(t, "エラー", o.Message)
			return nil
		}

		service.NewSearchService(s, &n, &testRandomizerMock{}, &loggerMock{}).NoticeError(&search.Request{
			Target: "12345",
		}, "エラー")
	}()
}
