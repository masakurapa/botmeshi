package service

import (
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/api"
	"github.com/masakurapa/botmeshi/app/domain/model/notification"
	"github.com/masakurapa/botmeshi/app/domain/repository"
	"github.com/masakurapa/botmeshi/app/domain/service"
	"github.com/masakurapa/botmeshi/test/mock"
	"github.com/stretchr/testify/assert"
)

type testNotificationMock struct {
	repository.Notification
	postMessageMock     func(notification.Option) error
	postRichMessageMock func(notification.Option) error
}

func (t *testNotificationMock) PostMessage(opt notification.Option) error {
	return t.postMessageMock(opt)
}

func (t *testNotificationMock) PostRichMessage(opt notification.Option) error {
	return t.postRichMessageMock(opt)
}

func TestNewEventService(t *testing.T) {
	func() {
		s := service.NewEventService(&testNotificationMock{}, mock.Logger())
		_, ok := s.(service.EventService)
		assert.True(t, ok)
	}()
}

func TestEventService_Exec(t *testing.T) {
	n := testNotificationMock{}
	text := "<1234567890> "

	// テキスト未送信
	func(n testNotificationMock) {
		n.postMessageMock = func(opt notification.Option) error {
			assert.Equal(t, "12345", opt.Target)
			assert.Equal(t, "探したい駅名を一緒に送って", opt.Message)
			return nil
		}
		n.postRichMessageMock = func(opt notification.Option) error {
			assert.Fail(t, "呼ばれないはず")
			return fmt.Errorf("post rich message error")
		}

		err := service.NewEventService(&n, mock.Logger()).Exec(&api.Parameter{
			ChannelID: "12345",
			Event: api.EventParameter{
				Text: "",
			},
		})
		assert.Nil(t, err)
	}(n)

	// 先頭12文字以外の文字列がない
	func(n testNotificationMock) {
		n.postMessageMock = func(opt notification.Option) error {
			assert.Equal(t, "12345", opt.Target)
			assert.Equal(t, "探したい駅名を一緒に送って", opt.Message)
			return nil
		}
		n.postRichMessageMock = func(opt notification.Option) error {
			assert.Fail(t, "呼ばれないはず")
			return fmt.Errorf("post rich message error")
		}

		err := service.NewEventService(&n, mock.Logger()).Exec(&api.Parameter{
			ChannelID: "12345",
			Event: api.EventParameter{
				Text: text,
			},
		})
		assert.Nil(t, err)
	}(n)

	// 先頭12文字以外の文字列がない
	func(n testNotificationMock) {
		n.postMessageMock = func(opt notification.Option) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		n.postRichMessageMock = func(opt notification.Option) error {
			assert.Equal(t, "12345", opt.Target)
			assert.Equal(t, "hoge で何が食べたい？", opt.Message)
			assert.Equal(t, "menu", opt.MessageID)
			assert.Equal(t, "#ff6633", opt.Color)
			assert.Equal(t, []notification.RichMessageOption{
				{
					ActionName: notification.ActionNameSelect,
					ActionType: notification.ActionTypeSelect,
					SelectOptions: []notification.SelectOption{
						{Text: "ラーメン", Value: "hoge ラーメン"},
						{Text: "肉", Value: "hoge 肉"},
						{Text: "魚", Value: "hoge 魚"},
						{Text: "定食", Value: "hoge 定食"},
						{Text: "カレー", Value: "hoge カレー"},
						{Text: "和食", Value: "hoge 和食"},
						{Text: "中華", Value: "hoge 中華"},
					},
				},
				{
					ActionName: notification.ActionNameCancel,
					ActionType: notification.ActionTypeButton,
					Text:       "やめる",
					Style:      "danger",
				},
			}, opt.RichMessageOptions)
			return nil
		}

		err := service.NewEventService(&n, mock.Logger()).Exec(&api.Parameter{
			ChannelID: "12345",
			Event: api.EventParameter{
				Text: text + "hoge",
			},
		})
		assert.Nil(t, err)
	}(n)

	// メッセージ送信エラー
	func(n testNotificationMock) {
		n.postMessageMock = func(opt notification.Option) error {
			assert.Fail(t, "呼ばれないはず")
			return nil
		}
		n.postRichMessageMock = func(opt notification.Option) error {
			return fmt.Errorf("post rich message error")
		}

		err := service.NewEventService(&n, mock.Logger()).Exec(&api.Parameter{
			ChannelID: "12345",
			Event: api.EventParameter{
				Text: text + "hoge",
			},
		})
		assert.NotNil(t, err)
		assert.Equal(t, "post rich message error", err.Error())
	}(n)
}
