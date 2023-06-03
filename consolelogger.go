package logger

import (
	"fmt"
)

type ConsoleLogger struct {
	baseLogger
	timeFormat string
}

func NewConsoleLogger() *ConsoleLogger {
	logger := &ConsoleLogger{
		timeFormat: TIMEFORMAT,
	}
	logger.baseLogger.log = logger.Log
	return logger
}

func (c *ConsoleLogger) Log(msg *Message) {
	fmt.Print(stdFormat(msg, c.timeFormat))
}
