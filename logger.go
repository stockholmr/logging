package logging

import (
	"fmt"
	"strings"
)

const (
	DEBUG = iota
	INFO
	REQUEST
	ERROR
	FATAL

	TIMEFORMAT = "2006-01-02 15:04:05"
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
	defaultLog = NewConsoleLogger(INFO)
}

func SetLogger(l Logger) {
	defaultLog = l
}

func Error(v ...interface{}) {
	defaultLog.Error(fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	defaultLog.Info(fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	defaultLog.Debug(fmt.Sprint(v...))
}

func Request(v ...interface{}) {
	defaultLog.Request(fmt.Sprint(v...))
}

func Fatal(v ...interface{}) {
	defaultLog.Fatal(fmt.Sprint(v...))
}

func stdFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat)
	buf += " [" + getLevelName(msg.level) + "]"
	buf += " " + msg.message

	if len(msg.message) > 0 && msg.message[len(msg.message)-1] != '\n' {
		buf += "\n"
	}
	return buf
}

func csvFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat) + ","
	buf += getLevelName(msg.level) + ","
	for i, o := range msg.csvMessage {
		v := fmt.Sprint(o)
		v = strings.Replace(strings.TrimSpace(v), ",", "_", -1)
		if i < len(msg.csvMessage)-1 {
			buf += fmt.Sprint(v) + ","
		} else {
			buf += fmt.Sprint(v)
		}
	}
	buf += "\n"
	return buf
}

func getLevelName(level int) string {
	switch level {
	case REQUEST:
		return "REQUEST"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}
