package logger

import (
	"os"
	"path"
	"testing"
)

var testFile string

func testfile() string {
	testFile = path.Join(os.TempDir(), "test.log")
	return testFile
}

func cleanup() {
	os.Remove(testFile)
}

func TestFileLogger(t *testing.T) {
	file, err := os.OpenFile(testfile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		t.Error(err)
	}

	logger := NewFileLogger(FORMAT_STD)
	logger.Open(file)
	logger.Info("TEST")
	logger.Close()

	cleanup()
}

func TestFileLogger_TestStopStart(t *testing.T) {
	file, err := os.OpenFile(testfile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		t.Error(err)
	}

	logger := NewFileLogger(FORMAT_STD)
	logger.Open(file)
	logger.Info("TEST")
	logger.Close()

	logger.Open(file)
	logger.Info("TEST")
	logger.Close()

	cleanup()
}
