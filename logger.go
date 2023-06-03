package logging

import (
	"fmt"
	"strings"
)

const (
	ERROR      string = "ERROR"
	INFO       string = "INFO"
	DEBUG      string = "DEBUG"
	FATAL      string = "FATAL"
	REQUEST    string = "REQUEST"
	TIMEFORMAT        = "2006-01-02 15:04:05"
)

var defaultLog Logger

type Logger interface {
	Error(...interface{})
	Info(...interface{})
	Debug(...interface{})
	Fatal(...interface{})
	Request(...interface{})
}

func init() {
	defaultLog = NewConsoleLogger()
}

func SetLogger(l Logger) {
	defaultLog = l
}

func Error(v ...interface{}) {
	defaultLog.Error(NewMessage(ERROR, fmt.Sprint(v...)))
}

func Info(v ...interface{}) {
	defaultLog.Info(fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	defaultLog.Debug(NewMessage(DEBUG, fmt.Sprint(v...)))
}

func Request(v ...interface{}) {
	defaultLog.Request(NewMessage(REQUEST, fmt.Sprint(v...)))
}

func Fatal(v ...interface{}) {
	defaultLog.Fatal(NewMessage(FATAL, fmt.Sprint(v...)))
}

func stdFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat)
	buf += " [" + msg.level + "]"
	buf += " " + msg.message

	if len(msg.message) > 0 && msg.message[len(msg.message)-1] != '\n' {
		buf += "\n"
	}
	return buf
}

func csvFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat) + ","
	buf += msg.level + ","
	buf += strings.Replace(strings.TrimSpace(msg.message), ",", "_", -1)
	buf += "\n"
	return buf
}
