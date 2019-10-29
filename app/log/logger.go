package log

import "fmt"

// Logger struct
type Logger interface {
	Start(string, string, ...interface{})
	End(string, string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
}

type logger struct {
}

// NewLogger returns Logger instance
func NewLogger() Logger {
	return &logger{}
}

// Start func log
func (l *logger) Start(cls, fnc string, args ...interface{}) {
	l.Info("Start "+cls+"."+fnc+"() parameters: ", args...)
}

// End func log
func (l *logger) End(cls, fnc string, args ...interface{}) {
	l.Info("End "+cls+"."+fnc+"() returns: ", args...)
}

// Info log
func (l *logger) Info(msg string, args ...interface{}) {
	l.log("INFO", msg, args...)
}

// Error log
func (l *logger) Error(msg string, args ...interface{}) {
	l.log("ERROR", msg, args...)
}

func (l *logger) log(level, msg string, args ...interface{}) {
	format := "[" + level + "]" + msg + ". "

	if len(args) == 0 {
		fmt.Println(format)
	} else {
		fmt.Printf(format+l.verb(args)+"\n", args...)
	}
}

func (*logger) verb(args ...interface{}) string {
	verb := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			verb += ", "
		}
		verb += "%#v"
	}
	return verb
}
