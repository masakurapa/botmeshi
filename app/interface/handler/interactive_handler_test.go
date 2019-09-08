package handler

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/usecase"
	"github.com/stretchr/testify/assert"
)

type testInteractiveUseCaseMock struct {
	usecase.UseCase
	parseMock    func(body string) (*api.Parameter, error)
	validateMock func() error
	execMock     func() error
}

func TestNewInteractiveHandler(t *testing.T) {
	func() {
		s := NewInteractiveHandler(&testInteractiveUseCaseMock{})
		_, ok := s.(Handler)
		assert.True(t, ok)
	}()
}

func (t *testInteractiveUseCaseMock) Parse(body string) (*api.Parameter, error) {
	return t.parseMock(body)
}
func (t *testInteractiveUseCaseMock) Validate(_ *api.Parameter) error {
	return t.validateMock()
}
func (t *testInteractiveUseCaseMock) Exec(_ *api.Parameter) error {
	return t.execMock()
}

func TestInteractiveHandler(t *testing.T) {
	uc := testInteractiveUseCaseMock{
		parseMock:    func(body string) (*api.Parameter, error) { return &api.Parameter{}, nil },
		validateMock: func() error { return nil },
		execMock:     func() error { return nil },
	}

	// 正常系
	func(uc testInteractiveUseCaseMock) {
		res, err := NewInteractiveHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "Success Interactive", res.Body)
		}
	}(uc)

	// Parse()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.parseMock = func(body string) (*api.Parameter, error) { return nil, fmt.Errorf("parse error") }

		res, err := NewInteractiveHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "parse error", res.Body)
		}
	}(uc)

	// Validate()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.validateMock = func() error { return fmt.Errorf("validate error") }

		res, err := NewInteractiveHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "validate error", res.Body)
		}
	}(uc)

	// Exec()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.execMock = func() error { return fmt.Errorf("exec error") }

		res, err := NewInteractiveHandler(&uc).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "exec error", res.Body)
		}
	}(uc)
}
