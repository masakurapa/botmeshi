package handler

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/usecase"
	"github.com/stretchr/testify/assert"
)

type testEventUseCaseMock struct {
	usecase.EventUseCase
	parseMock    func(body string) (*api.Parameter, error)
	validateMock func() error
	execMock     func() error
}

func (t *testEventUseCaseMock) Parse(body string) (*api.Parameter, error) {
	return t.parseMock(body)
}
func (t *testEventUseCaseMock) Validate(_ *api.Parameter) error {
	return t.validateMock()
}
func (t *testEventUseCaseMock) Exec(_ *api.Parameter) error {
	return t.execMock()
}

func TestNewEventHandler(t *testing.T) {
	func() {
		s := NewEventHandler(&testEventUseCaseMock{})
		_, ok := s.(Handler)
		assert.True(t, ok)
	}()
}

func TestEventHandler(t *testing.T) {
	uc := testEventUseCaseMock{
		parseMock:    func(body string) (*api.Parameter, error) { return &api.Parameter{}, nil },
		validateMock: func() error { return nil },
		execMock:     func() error { return nil },
	}

	// 正常系
	func(uc testEventUseCaseMock) {
		res, err := NewEventHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "Success Event", res.Body)
		}
	}(uc)

	// Parse()のエラー
	func(uc testEventUseCaseMock) {
		uc.parseMock = func(body string) (*api.Parameter, error) { return nil, fmt.Errorf("parse error") }

		res, err := NewEventHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "parse error", res.Body)
		}
	}(uc)

	// Validate()のエラー
	func(uc testEventUseCaseMock) {
		uc.validateMock = func() error { return fmt.Errorf("validate error") }

		res, err := NewEventHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "validate error", res.Body)
		}
	}(uc)

	// Exec()のエラー
	func(uc testEventUseCaseMock) {
		uc.execMock = func() error { return fmt.Errorf("exec error") }

		res, err := NewEventHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "exec error", res.Body)
		}
	}(uc)
}
