package notification_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/masakurapa/botmeshi/common/notification"
	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
	"github.com/masakurapa/botmeshi/common/test/testhelper"
)

type mockClient struct {
	notification.Client
	mockPostTextMessage        func(vo.Target, vo.Message) error
	mockPostInteractiveMessage func(vo.Target, vo.Message, attachment.Attachment) error
}

func (m *mockClient) PostTextMessage(vt vo.Target, vm vo.Message) error {
	return m.mockPostTextMessage(vt, vm)
}
func (m *mockClient) PostInteractiveMessage(vt vo.Target, vm vo.Message, att attachment.Attachment) error {
	return m.mockPostInteractiveMessage(vt, vm, att)
}

func TestNotification_TextMessage(t *testing.T) {
	tgt := testhelper.CreateTargetVO(t, "target")
	msg := testhelper.CreateMessageVO(t, "message")

	var actTgt vo.Target
	var actMsg vo.Message

	c := notification.New(&mockClient{
		mockPostTextMessage: func(vt vo.Target, vm vo.Message) error {
			actTgt = vt
			actMsg = vm
			return nil
		},
	}).TextMessage(tgt, msg)

	t.Run("Post()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err := c.Post(); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("PostTextMessage()のvo.TargetはTextMessage()の引数と同じ", func(t *testing.T) {
		if !reflect.DeepEqual(tgt, actTgt) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
	t.Run("PostTextMessage()のvo.MessageはTextMessage()の引数と同じ", func(t *testing.T) {
		if !reflect.DeepEqual(msg, actMsg) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})

	t.Run("Post()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c := notification.New(&mockClient{
			mockPostTextMessage: func(vt vo.Target, vm vo.Message) error {
				return mockError
			},
		}).TextMessage(tgt, msg)
		if err := c.Post(); err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}

func TestNotification_InteractiveSelectMessage(t *testing.T) {
	tgt := testhelper.CreateTargetVO(t, "target")
	msg := testhelper.CreateMessageVO(t, "message")
	tvs := testhelper.CreateTextValuesVO(t, testhelper.CreateTextValueVO(t, "t", "v"))

	var actTgt vo.Target
	var actMsg vo.Message
	var actAtt attachment.Attachment

	c := notification.New(&mockClient{
		mockPostInteractiveMessage: func(vt vo.Target, vm vo.Message, att attachment.Attachment) error {
			actTgt = vt
			actMsg = vm
			actAtt = att
			return nil
		},
	}).InteractiveSelectMessage(tgt, msg, tvs)

	t.Run("Post()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err := c.Post(); err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("PostInteractiveMessage()のvo.TargetはInteractiveSelectMessage()の引数と同じ", func(t *testing.T) {
		if !reflect.DeepEqual(tgt, actTgt) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
	t.Run("PostInteractiveMessage()のvo.MessageはInteractiveSelectMessage()の引数と同じ", func(t *testing.T) {
		if !reflect.DeepEqual(msg, actMsg) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})

	t.Run("PostInteractiveMessage()のattachment.Attachment().SelectBox()の長さはTextValuesと同じ", func(t *testing.T) {
		if len(tvs.All()) != len(actAtt.SelectBox().Options()) {
			t.Fatalf("want %d but got %d", len(tvs.All()),len(actAtt.SelectBox().Options()))
		}
	})
	t.Run("PostInteractiveMessage()のattachment.Attachment().Button()はattachment.CancelButtonが設定されている", func(t *testing.T) {
		if a := attachment.NewCancelButton().ActionName(); a.Name() != actAtt.Button().ActionName().Name() {
			t.Fatalf("want %s but got %s", a.Name(), actAtt.Button().ActionName().Name())
		}
	})

	t.Run("Post()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c := notification.New(&mockClient{
			mockPostInteractiveMessage: func(vt vo.Target, vm vo.Message, att attachment.Attachment) error {
				return mockError
			},
		}).InteractiveSelectMessage(tgt, msg, tvs)
		if err := c.Post(); err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}
