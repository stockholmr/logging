package logging

import "time"

type Message struct {
	level      int
	message    string
	time       time.Time
	csvMessage []interface{}
}

func NewMessage(level int, message string) *Message {
	return &Message{
		level:   level,
		message: message,
		time:    time.Now(),
	}
}

func NewCSVMessage(level int, message ...interface{}) *Message {
	return &Message{
		level:      level,
		csvMessage: message,
		time:       time.Now(),
	}
}
