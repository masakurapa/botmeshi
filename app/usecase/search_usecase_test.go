package usecase

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/stretchr/testify/assert"
)

type testSearchServiceMock struct {
	service.SearchService
	searchStationMock func() *search.Station
	searchShopsMock   func() []search.Shop
	noticeSuccessMock func(*search.Request, []search.Shop) error
	noticeErrorMock   func(*search.Request, string)
}

func (t *testSearchServiceMock) SearchStation(*search.Query) *search.Station {
	return t.searchStationMock()
}
func (t *testSearchServiceMock) SearchShops(*search.Query, *search.Station) []search.Shop {
	return t.searchShopsMock()
}
func (t *testSearchServiceMock) NoticeSuccess(p *search.Request, shops []search.Shop) error {
	return t.noticeSuccessMock(p, shops)
}
func (t *testSearchServiceMock) NoticeError(p *search.Request, s string) {
	t.noticeErrorMock(p, s)
}

func TestNewSearchUseCase(t *testing.T) {
	func() {
		s := NewSearchUseCase(&testSearchServiceMock{})
		_, ok := s.(SearchUseCase)
		assert.True(t, ok)
	}()
}

func TestSearchUseCase_Validate(t *testing.T) {
	s := &testSearchServiceMock{}

	// 正常系
	func() {
		err := NewSearchUseCase(s).Validate(&search.Request{
			Target: "hoge",
			Query:  "fuga",
		})
		assert.Nil(t, err)
	}()

	// Target空
	func() {
		err := NewSearchUseCase(s).Validate(&search.Request{
			Target: "",
			Query:  "fuga",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "target is required", err.Error())
	}()

	// Query空
	func() {
		err := NewSearchUseCase(s).Validate(&search.Request{
			Target: "hoge",
			Query:  "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "query is required", err.Error())
	}()
}

func TestSearchUseCase_Exec(t *testing.T) {
	s := testSearchServiceMock{}

	// 正常系
	func(s testSearchServiceMock) {
		s.searchStationMock = func() *search.Station { return &search.Station{} }
		s.searchShopsMock = func() []search.Shop { return []search.Shop{{}} }
		s.noticeSuccessMock = func(*search.Request, []search.Shop) error { return nil }
		s.noticeErrorMock = func(*search.Request, string) { assert.Fail(t, "呼ばれないはず") }

		err := NewSearchUseCase(&s).Exec(&search.Request{
			Query: "hoge fuga",
		})
		assert.Nil(t, err)
	}(s)

	// 正常系（クエリにスペース複数
	func(s testSearchServiceMock) {
		s.searchStationMock = func() *search.Station { return &search.Station{} }
		s.searchShopsMock = func() []search.Shop { return []search.Shop{{}} }
		s.noticeSuccessMock = func(*search.Request, []search.Shop) error { return nil }
		s.noticeErrorMock = func(*search.Request, string) { assert.Fail(t, "呼ばれないはず") }

		err := NewSearchUseCase(&s).Exec(&search.Request{
			Query: "hoge fuga hoga",
		})
		assert.Nil(t, err)
	}(s)

	// 異常系（クエリにスペースなし
	func(s testSearchServiceMock) {
		s.searchStationMock = func() *search.Station {
			assert.Fail(t, "呼ばれないはず")
			return &search.Station{}
		}
		s.searchShopsMock = func() []search.Shop {
			assert.Fail(t, "呼ばれないはず")
			return []search.Shop{{}}
		}
		s.noticeSuccessMock = func(*search.Request, []search.Shop) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s.noticeErrorMock = func(*search.Request, string) { assert.Fail(t, "呼ばれないはず") }

		err := NewSearchUseCase(&s).Exec(&search.Request{
			Query: "hoge",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid query: hoge", err.Error())
	}(s)

	// 異常系（店舗データ0件
	func(s testSearchServiceMock) {
		request := &search.Request{Query: "hoge fuga"}

		s.searchStationMock = func() *search.Station { return &search.Station{} }
		s.searchShopsMock = func() []search.Shop { return []search.Shop{} }
		s.noticeSuccessMock = func(*search.Request, []search.Shop) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s.noticeErrorMock = func(p *search.Request, s string) {
			assert.Equal(t, "お店が見つからなかったよ", s)
			assert.Equal(t, request, p)
		}

		err := NewSearchUseCase(&s).Exec(request)
		assert.NotNil(t, err)
		assert.Equal(t, "shop not found", err.Error())
	}(s)

	// 異常系（成功通知の失敗
	func(s testSearchServiceMock) {
		request := &search.Request{Query: "hoge fuga"}

		s.searchStationMock = func() *search.Station { return &search.Station{} }
		s.searchShopsMock = func() []search.Shop { return []search.Shop{{}} }
		s.noticeSuccessMock = func(*search.Request, []search.Shop) error { return fmt.Errorf("notice success error") }
		s.noticeErrorMock = func(p *search.Request, s string) {
			assert.Equal(t, "お店の通知ができなかったよ", s)
			assert.Equal(t, request, p)
		}

		err := NewSearchUseCase(&s).Exec(request)
		assert.NotNil(t, err)
		assert.Equal(t, "notice error", err.Error())
	}(s)
}
