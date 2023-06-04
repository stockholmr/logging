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
		b.llog(NewCSVMessage(ERROR, v...))
		return
	}
	b.llog(NewMessage(ERROR, fmt.Sprint(v...)))
}

func (b *baseLogger) Info(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(INFO, v...))
		return
	}
	b.llog(NewMessage(INFO, fmt.Sprint(v...)))
}

func (b *baseLogger) Debug(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(DEBUG, v...))
		return
	}
	b.llog(NewMessage(DEBUG, fmt.Sprint(v...)))
}

func (b *baseLogger) Request(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(REQUEST, v...))
		return
	}
	b.llog(NewMessage(REQUEST, fmt.Sprint(v...)))
}

func (b *baseLogger) Fatal(v ...interface{}) {
	if b.useCSVFormat {
		b.llog(NewCSVMessage(FATAL, v...))
		os.Exit(1)
		return
	}
	b.llog(NewMessage(FATAL, fmt.Sprint(v...)))
	os.Exit(1)
}
