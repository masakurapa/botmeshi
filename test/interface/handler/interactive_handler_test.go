package handler

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/masakurapa/botmeshi/app/interface/handler"
	"github.com/masakurapa/botmeshi/app/usecase"
	"github.com/masakurapa/botmeshi/test/mock"
	"github.com/stretchr/testify/assert"
)

type testInteractiveUseCaseMock struct {
	usecase.InteractiveUseCase
	parseMock    func(body string) (*api.Parameter, error)
	validateMock func() error
	execMock     func() (string, error)
}

func (t *testInteractiveUseCaseMock) Parse(body string) (*api.Parameter, error) {
	return t.parseMock(body)
}
func (t *testInteractiveUseCaseMock) Validate(_ *api.Parameter) error {
	return t.validateMock()
}
func (t *testInteractiveUseCaseMock) Exec(_ *api.Parameter) (string, error) {
	return t.execMock()
}

func TestNewInteractiveHandler(t *testing.T) {
	func() {
		s := handler.NewInteractiveHandler(&testInteractiveUseCaseMock{}, mock.Logger())
		_, ok := s.(handler.Handler)
		assert.True(t, ok)
	}()
}

func TestInteractiveHandler(t *testing.T) {
	uc := testInteractiveUseCaseMock{
		parseMock:    func(body string) (*api.Parameter, error) { return &api.Parameter{}, nil },
		validateMock: func() error { return nil },
		execMock:     func() (string, error) { return "exec success", nil },
	}

	// 正常系
	func(uc testInteractiveUseCaseMock) {
		res, err := handler.NewInteractiveHandler(&uc, mock.Logger()).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "exec success", res.Body)
		}
	}(uc)

	// Parse()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.parseMock = func(body string) (*api.Parameter, error) { return nil, fmt.Errorf("parse error") }

		res, err := handler.NewInteractiveHandler(&uc, mock.Logger()).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "parse error", res.Body)
		}
	}(uc)

	// Validate()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.validateMock = func() error { return fmt.Errorf("validate error") }

		res, err := handler.NewInteractiveHandler(&uc, mock.Logger()).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "validate error", res.Body)
		}
	}(uc)

	// Exec()のエラー
	func(uc testInteractiveUseCaseMock) {
		uc.execMock = func() (string, error) { return "exec success", fmt.Errorf("exec error") }

		res, err := handler.NewInteractiveHandler(&uc, mock.Logger()).Handler(http.Request{})
		if assert.Nil(t, err) {
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, "exec error", res.Body)
		}
	}(uc)
}
