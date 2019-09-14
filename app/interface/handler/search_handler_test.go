package handler

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/search"
	"github.com/masakurapa/botmeshi/app/usecase"
	"github.com/stretchr/testify/assert"
)

type testSearchUseCaseMock struct {
	usecase.SearchUseCase
	validateMock func() error
	execMock     func() error
}

func (t *testSearchUseCaseMock) Validate(p *search.Request) error {
	return t.validateMock()
}
func (t *testSearchUseCaseMock) Exec(p *search.Request) error {
	return t.execMock()
}

func TestNewSearchHandler(t *testing.T) {
	func() {
		s := NewSearchHandler(&testSearchUseCaseMock{})
		_, ok := s.(SearchHandler)
		assert.True(t, ok)
	}()
}

func TestSearchHandler(t *testing.T) {
	uc := testSearchUseCaseMock{}
	p := search.Request{Target: "", Query: ""}

	// 正常系
	func(uc testSearchUseCaseMock) {
		uc.validateMock = func() error { return nil }
		uc.execMock = func() error { return nil }

		msg, err := NewSearchHandler(&uc).Handler(p)
		assert.Nil(t, err)
		assert.Equal(t, "Success Search", msg)
	}(uc)

	// バリデーションエラー
	func(uc testSearchUseCaseMock) {
		uc.validateMock = func() error { return fmt.Errorf("validate error") }
		msg, err := NewSearchHandler(&uc).Handler(p)
		assert.Nil(t, err)
		assert.Equal(t, "validate error", msg)
	}(uc)

	// 実行エラー
	func(uc testSearchUseCaseMock) {
		uc.validateMock = func() error { return nil }
		uc.execMock = func() error { return fmt.Errorf("exec error") }
		msg, err := NewSearchHandler(&uc).Handler(p)
		assert.Nil(t, err)
		assert.Equal(t, "exec error", msg)
	}(uc)
}
