package testhelper

import (
	"testing"

	"github.com/masakurapa/botmeshi/common/notification/domain/vo"
)

// CreateMessageVO はvo.Messageを生成して返す
func CreateMessageVO(t *testing.T, message string) vo.Message {
	t.Helper()
	m, err := vo.NewMessage(message)
	if err != nil {
		t.Fatalf("vo.NewMessage() returned %q", err.Error())
	}
	return m
}

// CreateTargetVO はvo.Targetを生成して返す
func CreateTargetVO(t *testing.T, target string) vo.Target {
	t.Helper()
	tg, err := vo.NewTarget(target)
	if err != nil {
		t.Fatalf("vo.NewTarget() returned %q", err.Error())
	}
	return tg
}

// CreateTextValueVO はvo.TextValueを生成して返す
func CreateTextValueVO(t *testing.T, text, value string) vo.TextValue {
	t.Helper()
	tv, err := vo.NewTextValue(text, value)
	if err != nil {
		t.Fatalf("vo.NewTextValue() returned %q", err.Error())
	}
	return tv
}

// CreateTextValuesVO はvo.TextValuesを生成して返す
func CreateTextValuesVO(t *testing.T, tv ...vo.TextValue) vo.TextValues {
	t.Helper()
	tvs, err := vo.NewTextValues(tv)
	if err != nil {
		t.Fatalf("vo.NewTextValues() returned %q", err.Error())
	}
	return tvs
}
