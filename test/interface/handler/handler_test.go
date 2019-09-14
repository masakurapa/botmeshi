package handler

import (
	"github.com/masakurapa/botmeshi/app/log"
)

type loggerMock struct {
	log.Logger
}

func (*loggerMock) Info(string, ...interface{})  {}
func (*loggerMock) Error(string, ...interface{}) {}
