package service

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/invoke"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/stretchr/testify/assert"
)

type testInfokeFunctionMock struct {
	repository.InvokeFunction
	execMock func(*invoke.Parameter) error
}

func (t *testInfokeFunctionMock) Exec(p *invoke.Parameter) error {
	return t.execMock(p)
}

func TestNewInteractiveService(t *testing.T) {
	func() {
		s := NewInteractiveService(&testInfokeFunctionMock{})
		_, ok := s.(InteractiveService)
		assert.True(t, ok)
	}()
}

func TestInvokeService_Exec(t *testing.T) {
	fnc := testInfokeFunctionMock{}

	// キャンセルアクション
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            notification.ActionNameCancel,
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "ばいびー", s)
	}(fnc)

	// ゴールアクション
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            notification.ActionNameGo,
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "`fuga` に\nごーーーーーーーーーーーーーーーる！！", s)
	}(fnc)

	// 非ゴールアクション
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            notification.ActionNameDoNotGo,
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "ざんねん...。", s)
	}(fnc)

	// 選択アクション
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			assert.Equal(t, "12345", p.Target)
			assert.Equal(t, "fuga", p.Query)
			return nil
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            notification.ActionNameSelect,
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "`fuga` でお店を探すよ！\nちょっと時間がかかるからまってくれ！", s)
	}(fnc)

	// 選択アクション（エラーあり）
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			return fmt.Errorf("exec error")
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            notification.ActionNameSelect,
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.NotNil(t, err)
		assert.Equal(t, "exec error", err.Error())
		assert.Equal(t, "", s)
	}(fnc)

	// 不正なアクション
	func(fnc testInfokeFunctionMock) {
		fnc.execMock = func(p *invoke.Parameter) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		s, err := NewInteractiveService(&fnc).Exec(&api.Parameter{
			ChannelID: "12345",
			Action: api.ActionParameter{
				Name:            "hoge",
				SelectedOptions: []string{"fuga", "fuge"},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "キサマ何者だ！", s)
	}(fnc)
}
