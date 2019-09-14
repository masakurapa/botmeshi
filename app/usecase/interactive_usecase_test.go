package usecase

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/stretchr/testify/assert"
)

type testInteractiveServiceMock struct {
	service.InteractiveService
	execMock func() (string, error)
}

func (t *testInteractiveServiceMock) Exec(p *api.Parameter) (string, error) {
	return t.execMock()
}

func TestNewInteractiveUseCase(t *testing.T) {
	func() {
		s := NewInteractiveUseCase(&testInteractiveServiceMock{}, &loggerMock{})
		_, ok := s.(InteractiveUseCase)
		assert.True(t, ok)
	}()
}

func TestInteractiveUseCase_Parse(t *testing.T) {
	s := testInteractiveServiceMock{}

	// 正常系
	func() {
		p, err := NewInteractiveUseCase(&s, &loggerMock{}).Parse("payload=%7B%7D")
		if assert.Nil(t, err) {
			assert.Equal(t, reflect.TypeOf(&api.Parameter{}), reflect.TypeOf(p))
		}
	}()

	// 異常系（不正なURLエンコード文字列
	func() {
		p, err := NewInteractiveUseCase(&s, &loggerMock{}).Parse("payload=%")
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}()

	// 異常系（不正なJSON
	func() {
		p, err := NewInteractiveUseCase(&s, &loggerMock{}).Parse("payload=%7B")
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}()
}

func TestInteractiveUseCase_Validate(t *testing.T) {
	s := testInteractiveServiceMock{}
	token := "interactive token"
	os.Setenv("BOT_VERIFICATION_TOKEN", token)

	// 正常系
	func() {
		err := NewInteractiveUseCase(&s, &loggerMock{}).Validate(&api.Parameter{Token: token})
		assert.Nil(t, err)
	}()

	// トークンエラー
	func() {
		err := NewInteractiveUseCase(&s, &loggerMock{}).Validate(&api.Parameter{Token: "error"})
		assert.NotNil(t, err)
		assert.Equal(t, "token error", err.Error())
	}()
}

func TestInteractiveUseCase_Exec(t *testing.T) {
	s := testInteractiveServiceMock{}

	// 正常系
	func(s testInteractiveServiceMock) {
		s.execMock = func() (string, error) { return "success", nil }
		actual, err := NewInteractiveUseCase(&s, &loggerMock{}).Exec(&api.Parameter{})
		assert.Nil(t, err)
		assert.Equal(t, "success", actual)
	}(s)

	// 異常系
	func(s testInteractiveServiceMock) {
		s.execMock = func() (string, error) { return "", fmt.Errorf("exec error") }
		actual, err := NewInteractiveUseCase(&s, &loggerMock{}).Exec(&api.Parameter{})
		assert.NotNil(t, err)
		assert.Equal(t, "exec error", err.Error())
		assert.Equal(t, "", actual)
	}(s)
}
