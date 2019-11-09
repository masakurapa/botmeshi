package slack_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/masakurapa/botmeshi/common/notification"
	"github.com/masakurapa/botmeshi/common/notification/domain/attachment"
	"github.com/masakurapa/botmeshi/common/notification/errs"
	ntc "github.com/masakurapa/botmeshi/common/notification/slack"
	"github.com/masakurapa/botmeshi/common/test/testhelper"
	slk "github.com/nlopes/slack"
)

type mockSendSlackClient struct {
	ntc.Client
	mockPostMessage func(string, ...slk.MsgOption) (string, string, error)
}

func (m *mockSendSlackClient) PostMessage(c string, opts ...slk.MsgOption) (string, string, error) {
	return m.mockPostMessage(c, opts...)
}

func TestNewSlackClient(t *testing.T) {
	n, err := ntc.NewClient("token")
	t.Run("正常終了時のエラーはnilを返す", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("notification.Clientの実装クラスを返す", func(t *testing.T) {
		if _, ok := n.(notification.Client); !ok {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
	t.Run("必須エラー時はerrs.RequiredArgsErrorの構造体を返す", func(t *testing.T) {
		if _, err := ntc.NewClient(""); errors.Is(err, &errs.RequiredArgsError{}) {
			t.Fatalf("want %t but got %t", true, false)
		}
	})
}

func TestPostTextMessage(t *testing.T) {
	vt := testhelper.CreateTargetVO(t, "target")
	vm := testhelper.CreateMessageVO(t, "message")

	var actTgt string
	var actOpts []slk.MsgOption

	c := ntc.SendSlackClient{
		Client: &mockSendSlackClient{
			mockPostMessage: func(c string, opts ...slk.MsgOption) (string, string, error) {
				actTgt = c
				actOpts = opts
				return "", "", nil
			},
		},
	}
	err := c.PostTextMessage(vt, vm)

	t.Run("Post()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("PostMessage()の第一引数はTarget.String()と等しい", func(t *testing.T) {
		if vt.String() != actTgt {
			t.Fatalf("want %q but got %q", vt.String(), actTgt)
		}
	})
	t.Run("PostMessage()の第二引数の長さは1と等しい", func(t *testing.T) {
		if 1 != len(actOpts) {
			t.Fatalf("want %d but got %d", 1, len(actOpts))
		}
	})

	t.Run("Post()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c := ntc.SendSlackClient{
			Client: &mockSendSlackClient{
				mockPostMessage: func(c string, opts ...slk.MsgOption) (string, string, error) {
					actTgt = c
					actOpts = opts
					return "", "", mockError
				},
			},
		}
		if err := c.PostTextMessage(vt, vm); err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}

func TestPostInteractiveMessage(t *testing.T) {
	vt := testhelper.CreateTargetVO(t, "target")
	vm := testhelper.CreateMessageVO(t, "message")

	tvs := testhelper.CreateTextValuesVO(t, testhelper.CreateTextValueVO(t, "t", "v"))
	att := attachment.NewAttachment(attachment.NewSelectBox(tvs), attachment.NewCancelButton())

	var actTgt string
	var actOpts []slk.MsgOption

	c := ntc.SendSlackClient{
		Client: &mockSendSlackClient{
			mockPostMessage: func(c string, opts ...slk.MsgOption) (string, string, error) {
				actTgt = c
				actOpts = opts
				return "", "", nil
			},
		},
	}
	err := c.PostInteractiveMessage(vt, vm, att)

	t.Run("Post()は正常終了時にエラーを返さない", func(t *testing.T) {
		if err != nil {
			t.Fatalf("want %v but got %q", nil, err.Error())
		}
	})
	t.Run("PostMessage()の第一引数はTarget.String()と等しい", func(t *testing.T) {
		if vt.String() != actTgt {
			t.Fatalf("want %q but got %q", vt.String(), actTgt)
		}
	})
	t.Run("PostMessage()の第二引数の長さは1と等しい", func(t *testing.T) {
		if 1 != len(actOpts) {
			t.Fatalf("want %d but got %d", 1, len(actOpts))
		}
	})

	t.Run("Post()は異常終了時にクライアントから受け取ったエラー返す", func(t *testing.T) {
		mockError := fmt.Errorf("post error")
		c := ntc.SendSlackClient{
			Client: &mockSendSlackClient{
				mockPostMessage: func(c string, opts ...slk.MsgOption) (string, string, error) {
					return "", "", mockError
				},
			},
		}
		if err := c.PostInteractiveMessage(vt, vm, att); err == nil {
			t.Fatalf("want %q but got %v", mockError.Error(), nil)
		}
	})
}
