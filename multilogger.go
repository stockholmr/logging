package logging

type MultiLogger struct {
	baseLogger
	loggers []Logger
}

func NewMultiLogger() *MultiLogger {
	logger := &MultiLogger{
		loggers: make([]Logger, 0),
	}
	logger.baseLogger.log = logger.Log
	return logger
}

func (m *MultiLogger) Add(logger Logger) {
	m.loggers = append(m.loggers, logger)
}

func (m *MultiLogger) SetLevel(level int) {
	for _, l := range m.loggers {
		l.SetLevel(level)
	}
}

func (m *MultiLogger) Log(msg *Message) {
	for _, l := range m.loggers {
		l.Log(msg)
	}
}
