package logging

import (
	"fmt"
)

type ConsoleLogger struct {
	baseLogger
	timeFormat string
	level      int
}

func NewConsoleLogger(level int) *ConsoleLogger {
	logger := &ConsoleLogger{
		timeFormat: TIMEFORMAT,
		level:      level,
	}
	logger.baseLogger.log = logger.Log
	return logger
}

func (c *ConsoleLogger) Log(msg *Message) {
	if msg.level >= c.level {
		fmt.Print(stdFormat(msg, c.timeFormat))
	}
}
