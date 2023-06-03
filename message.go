package logging

import "time"

type Message struct {
	level      string
	message    string
	time       time.Time
	csvMessage []interface{}
}

func NewMessage(level string, message string) *Message {
	return &Message{
		level:   level,
		message: message,
		time:    time.Now(),
	}
}

func NewCSVMessage(level string, message ...interface{}) *Message {
	return &Message{
		level:      level,
		csvMessage: message,
		time:       time.Now(),
	}
}
