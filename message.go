package logging

import "time"

type Message struct {
	level   string
	message string
	time    time.Time
}

func NewMessage(level string, message string) *Message {
	return &Message{
		level:   level,
		message: message,
		time:    time.Now(),
	}
}
