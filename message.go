package logging

import "time"

type Message struct {
	level   int
	message []interface{}
	time    time.Time
}

func NewMessage(level int, message ...interface{}) *Message {
	return &Message{
		level:   level,
		message: message,
		time:    time.Now(),
	}
}
