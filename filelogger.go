package logging

import (
	"log"
	"os"
)

const (
	BUFSIZE = 100

	FORMAT_STD = 0
	FORMAT_CSV = 1
)

type FileLogger struct {
	baseLogger
	queue      chan *Message
	done       chan bool
	closed     bool
	out        *os.File
	timeFormat string
	format     int
	level      int
}

func NewFileLogger(level int, format int) *FileLogger {
	logger := &FileLogger{
		closed:     false,
		timeFormat: TIMEFORMAT,
		format:     format,
		level:      level,
	}
	logger.baseLogger.log = logger.Log
	return logger
}

func (f *FileLogger) listen() {
	for {
		m, ok := <-f.queue
		if !ok {
			// the channel is closed and empty
			log.Print("Closing log file")
			_ = f.out.Sync()
			if err := f.out.Close(); err != nil {
				log.Print(err)
			}
			f.done <- true
			return
		}
		f.write(m)
	}
}

func (f *FileLogger) Open(file *os.File) {
	f.queue = make(chan *Message, BUFSIZE)
	f.done = make(chan bool)
	f.out = file

	go f.listen()
}

func (f *FileLogger) Close() {
	f.closed = true
	close(f.queue)
	<-f.done
}

func (f *FileLogger) SetLevel(level int) {
	f.level = level
}

func (f *FileLogger) Log(msg *Message) {
	if msg.level >= f.level {
		f.queue <- msg
		recover()
	}
}

func (f *FileLogger) write(msg *Message) {
	var buf string

	if f.format == FORMAT_STD {
		buf = stdFormat(msg, f.timeFormat)
	}

	if f.format == FORMAT_CSV {
		buf = csvFormat(msg, f.timeFormat)
	}

	_, _ = f.out.WriteString(buf)
}
