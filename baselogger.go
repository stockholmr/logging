package logging

import (
	"fmt"
	"os"
)

type LogFn func(*Message)

type baseLogger struct {
	log          LogFn
	useCSVFormat bool
}

func (b *baseLogger) llog(message *Message) {
	b.log(message)
}

func (b *baseLogger) Error(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(LOGLEVEL_ERROR, v...))
		return
	}
	b.llog(NewMessage(LOGLEVEL_ERROR, fmt.Sprint(v...)))
}

func (b *baseLogger) Info(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(LOGLEVEL_INFO, v...))
		return
	}
	b.llog(NewMessage(LOGLEVEL_INFO, fmt.Sprint(v...)))
}

func (b *baseLogger) Debug(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(LOGLEVEL_DEBUG, v...))
		return
	}
	b.llog(NewMessage(LOGLEVEL_DEBUG, fmt.Sprint(v...)))
}

func (b *baseLogger) Request(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(LOGLEVEL_REQUEST, v...))
		return
	}
	b.llog(NewMessage(LOGLEVEL_REQUEST, fmt.Sprint(v...)))
}

func (b *baseLogger) Fatal(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(LOGLEVEL_FATAL, v...))
		os.Exit(1)
		return
	}
	b.llog(NewMessage(LOGLEVEL_FATAL, fmt.Sprint(v...)))
	os.Exit(1)
}
