package logging

import (
	"fmt"
	"os"
)

type LogFn func(*Message)

type baseLogger struct {
	log LogFn
}

func (b *baseLogger) llog(message *Message) {
	b.log(message)
}

func (b *baseLogger) Error(v ...interface{}) {
	b.llog(NewMessage(ERROR, fmt.Sprint(v...)))
}

func (b *baseLogger) Info(v ...interface{}) {
	b.llog(NewMessage(INFO, fmt.Sprint(v...)))
}

func (b *baseLogger) Debug(v ...interface{}) {
	b.llog(NewMessage(DEBUG, fmt.Sprint(v...)))
}

func (b *baseLogger) Request(v ...interface{}) {
	b.llog(NewMessage(REQUEST, fmt.Sprint(v...)))
}

func (b *baseLogger) Fatal(v ...interface{}) {
	b.llog(NewMessage(FATAL, fmt.Sprint(v...)))
	os.Exit(1)
}
