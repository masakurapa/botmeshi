package usecase

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/app/util"
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
		s := NewEventUseCase(&testEventServiceMock{}, &loggerMock{})
		_, ok := s.(EventUseCase)
		assert.True(t, ok)
	}()
}

func TestEventUseCase_Parse(t *testing.T) {
	s := &testEventServiceMock{}

	// 正常系
	func() {
		p, err := NewEventUseCase(s, &loggerMock{}).Parse("{}")
		assert.Nil(t, err)
		if assert.NotNil(t, p) {
			assert.Equal(t, reflect.TypeOf(&api.Parameter{}), reflect.TypeOf(p))
		}
	}()

	// 異常系
	func() {
		p, err := NewEventUseCase(s, &loggerMock{}).Parse("not json body")
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}()
}

func TestEventUseCase_Validate(t *testing.T) {
	s := &testEventServiceMock{}
	token := "event token"
	os.Setenv(util.APIVerificationTokenKey, token)

	// 正常系
	func() {
		err := NewEventUseCase(s, &loggerMock{}).Validate(&api.Parameter{
			Token:     token,
			Type:      "event",
			Challenge: "challenge",
		})
		assert.Nil(t, err)
	}()

	// トークンエラー
	func() {
		err := NewEventUseCase(s, &loggerMock{}).Validate(&api.Parameter{Token: "error"})
		assert.NotNil(t, err)
		assert.Equal(t, "token error", err.Error())
	}()

	// URL検証
	func() {
		err := NewEventUseCase(s, &loggerMock{}).Validate(&api.Parameter{
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
		err := NewEventUseCase(&s, &loggerMock{}).Exec(&api.Parameter{})
		assert.Nil(t, err)
	}(s)

	// 異常系
	func(s testEventServiceMock) {
		s.execMock = func() error { return fmt.Errorf("exec error") }
		err := NewEventUseCase(&s, &loggerMock{}).Exec(&api.Parameter{})
		assert.NotNil(t, err)
		assert.Equal(t, "exec error", err.Error())
	}(s)
}
