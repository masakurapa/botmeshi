package usecase

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/usecase"
	"github.com/masakurapa/botmeshi/test/mock"
	"github.com/stretchr/testify/assert"
)

type testEventServiceMock struct {
	service.EventService
	execMock func() error
}

func (t *testEventServiceMock) Exec(p *api.Parameter) error {
	return t.execMock()
}

func TestNewEventUseCase(t *testing.T) {
	func() {
		s := usecase.NewEventUseCase(&testEventServiceMock{}, mock.Logger())
		_, ok := s.(usecase.EventUseCase)
		assert.True(t, ok)
	}()
}

func TestEventUseCase_Parse(t *testing.T) {
	s := &testEventServiceMock{}

	// 正常系
	func() {
		p, err := usecase.NewEventUseCase(s, mock.Logger()).Parse("{}")
		assert.Nil(t, err)
		if assert.NotNil(t, p) {
			assert.Equal(t, reflect.TypeOf(&api.Parameter{}), reflect.TypeOf(p))
		}
	}()

	// 異常系
	func() {
		p, err := usecase.NewEventUseCase(s, mock.Logger()).Parse("not json body")
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}()
}

func TestEventUseCase_Validate(t *testing.T) {
	s := &testEventServiceMock{}
	token := "event token"
	os.Setenv("BOT_VERIFICATION_TOKEN", token)

	// 正常系
	func() {
		err := usecase.NewEventUseCase(s, mock.Logger()).Validate(&api.Parameter{
			Token:     token,
			Type:      "event",
			Challenge: "challenge",
		})
		assert.Nil(t, err)
	}()

	// トークンエラー
	func() {
		err := usecase.NewEventUseCase(s, mock.Logger()).Validate(&api.Parameter{Token: "error"})
		assert.NotNil(t, err)
		assert.Equal(t, "token error", err.Error())
	}()

	// URL検証
	func() {
		err := usecase.NewEventUseCase(s, mock.Logger()).Validate(&api.Parameter{
			Token:     token,
			Type:      "url_verification",
			Challenge: "challenge",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "challenge", err.Error())
	}()
}

func TestEventUseCase_Exec(t *testing.T) {
	s := testEventServiceMock{}

	// 正常系
	func(s testEventServiceMock) {
		s.execMock = func() error { return nil }
		err := usecase.NewEventUseCase(&s, mock.Logger()).Exec(&api.Parameter{})
		assert.Nil(t, err)
	}(s)

	// 異常系
	func(s testEventServiceMock) {
		s.execMock = func() error { return fmt.Errorf("exec error") }
		err := usecase.NewEventUseCase(&s, mock.Logger()).Exec(&api.Parameter{})
		assert.NotNil(t, err)
		assert.Equal(t, "exec error", err.Error())
	}(s)
}
