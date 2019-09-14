package mock

import "github.com/masakurapa/botmeshi/app/log"

type loggerMock struct {
	log.Logger
}

// Logger ロガーの共通モックを返す
func Logger() log.Logger {
	return &loggerMock{}
}

func (*loggerMock) Start(string, string, ...interface{}) {}
func (*loggerMock) End(string, string, ...interface{})   {}
func (*loggerMock) Info(string, ...interface{})          {}
func (*loggerMock) Error(string, ...interface{})         {}
