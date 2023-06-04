package logging

import (
	"fmt"
	"strings"
)

const (
	DEBUG = iota
	REQUEST
	INFO
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
	Log(message *Message)
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

func joinSlice(s []interface{}, delimiter string) string {
	var str string
	for i, o := range s {
		v := strings.TrimSpace(fmt.Sprint(o))
		if i < len(s)-1 {
			str += v + delimiter
		} else {
			str += v
		}
	}
	return str
}

func stdFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat)
	buf += " [" + getLevelName(msg.level) + "]"
	buf += " " + joinSlice(msg.message, " ")

	if len(msg.message) > 0 && msg.message[len(msg.message)-1] != '\n' {
		buf += "\n"
	}
	return buf
}

func csvFormat(msg *Message, timeFormat string) string {
	buf := msg.time.Format(timeFormat) + ","
	buf += getLevelName(msg.level) + ","
	buf += joinSlice(msg.message, ",")
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
