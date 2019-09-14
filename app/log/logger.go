package log

import "fmt"

// Logger struct
type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
}

type logger struct {
}

// NewLogger returns Logger instance
func NewLogger() Logger {
	return &logger{}
}

// Info log
func (l *logger) Info(format string, arg ...interface{}) {
	l.log("INFO", format, arg...)
}

// Error log
func (l *logger) Error(format string, arg ...interface{}) {
	l.log("ERROR", format, arg...)
}

func (*logger) log(level, format string, arg ...interface{}) {
	fmt.Printf("["+level+"]"+format+"\n", arg...)
}
