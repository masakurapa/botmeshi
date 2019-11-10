package faashelper

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
)

// CreateFunctionNameVO はvo.FunctionNameを生成して返す
func CreateFunctionNameVO(t *testing.T, name string) vo.FunctionName {
	t.Helper()
	v, err := vo.NewFunctionName(name)
	if err != nil {
		t.Fatalf("vo.NewFunctionName() returned %q", err.Error())
	}
	return v
}

// CreateNotificationCallbackTargetVO はvo.NotificationCallbackTargetを生成して返す
func CreateNotificationCallbackTargetVO(t *testing.T, target string) vo.NotificationCallbackTarget {
	t.Helper()
	v, err := vo.NewNotificationCallbackTarget(target)
	if err != nil {
		t.Fatalf("vo.NewNotificationCallbackTarget() returned %q", err.Error())
	}
	return v
}

// CreatePayloadVO はvo.Payloadを生成して返す
func CreatePayloadVO(t *testing.T, payload string) vo.Payload {
	t.Helper()
	v, err := vo.NewPayload(payload)
	if err != nil {
		t.Fatalf("vo.NewPayload() returned %q", err.Error())
	}
	return v
}

// CreateSearchQueryVO はvo.SearchQueryを生成して返す
func CreateSearchQueryVO(t *testing.T, query string) vo.SearchQuery {
	t.Helper()
	v, err := vo.NewSearchQuery(query)
	if err != nil {
		t.Fatalf("vo.NewSearchQuery() returned %q", err.Error())
	}
	return v
}
