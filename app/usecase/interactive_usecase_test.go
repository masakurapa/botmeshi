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

type testInteractiveServiceMock struct {
	service.InteractiveService
	execMock func() (string, error)
}

func (t *testInteractiveServiceMock) Exec(p *api.Parameter) (string, error) {
	return t.execMock()
}

func TestNewInteractiveUseCase(t *testing.T) {
	func() {
		s := NewInteractiveUseCase(&testInteractiveServiceMock{})
		_, ok := s.(InteractiveUseCase)
		assert.True(t, ok)
	}()
}

func TestInteractiveUseCase_Parse(t *testing.T) {
	s := testInteractiveServiceMock{}

	// 正常系
	func() {
		p, err := NewInteractiveUseCase(&s).Parse("{}")
		assert.Nil(t, err)
		if assert.NotNil(t, p) {
			assert.Equal(t, reflect.TypeOf(&api.Parameter{}), reflect.TypeOf(p))
		}
	}()

	// 異常系
	func() {
		p, err := NewInteractiveUseCase(&s).Parse("not json body")
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}()
}

func TestInteractiveUseCase_Validate(t *testing.T) {
	s := testInteractiveServiceMock{}
	token := "interactive token"
	os.Setenv(util.APIVerificationTokenKey, token)

	// 正常系
	func() {
		err := NewInteractiveUseCase(&s).Validate(&api.Parameter{Token: token})
		assert.Nil(t, err)
	}()

	// トークンエラー
	func() {
		err := NewInteractiveUseCase(&s).Validate(&api.Parameter{Token: "error"})
		assert.NotNil(t, err)
		assert.Equal(t, "token error", err.Error())
	}()
}

func TestInteractiveUseCase_Exec(t *testing.T) {
	s := testInteractiveServiceMock{}

	// 正常系
	func(s testInteractiveServiceMock) {
		s.execMock = func() (string, error) { return "success", nil }
		actual, err := NewInteractiveUseCase(&s).Exec(&api.Parameter{})
		assert.Nil(t, err)
		assert.Equal(t, "success", actual)
	}(s)

	// 異常系
	func(s testInteractiveServiceMock) {
		s.execMock = func() (string, error) { return "", fmt.Errorf("exec error") }
		actual, err := NewInteractiveUseCase(&s).Exec(&api.Parameter{})
		assert.NotNil(t, err)
		assert.Equal(t, "exec error", err.Error())
		assert.Equal(t, "", actual)
	}(s)
}
