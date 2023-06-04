package logging

import (
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
	b.llog(NewMessage(ERROR, v...))
}

func (b *baseLogger) Info(v ...interface{}) {
	b.llog(NewMessage(INFO, v...))
}

func (b *baseLogger) Debug(v ...interface{}) {
	b.llog(NewMessage(DEBUG, v...))
}

func (b *baseLogger) Request(v ...interface{}) {
	b.llog(NewMessage(REQUEST, v...))
}

func (b *baseLogger) Fatal(v ...interface{}) {
	b.llog(NewMessage(FATAL, v...))
	os.Exit(1)
}
