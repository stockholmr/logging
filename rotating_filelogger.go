package logging

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type logFile struct {
	Filename  string
	LogNumber int
}

type RotatingFileLogger struct {
	baseLogger
	fileLogger    *FileLogger
	filePath      string
	fileName      string
	fileExt       string
	suffixPattern string
	maxFileCount  int
	lineCount     int64
	maxLines      int64
	level         int
}

func NewRotatingFileLogger(logDir string, logFilename string, fileExtension string, level int, format int, maxLines int64, maxFiles int) *RotatingFileLogger {
	logger := &RotatingFileLogger{
		fileLogger: &FileLogger{
			closed:     false,
			timeFormat: TIMEFORMAT,
			format:     format,
			level:      level,
		},

		filePath:     logDir,
		fileName:     logFilename,
		fileExt:      fileExtension,
		maxLines:     maxLines,
		maxFileCount: maxFiles,
		level:        level,
	}
	logger.baseLogger.log = logger.Log

	if logger.fileExt[0:1] != "." {
		logger.fileExt = "." + logger.fileExt
	}

	return logger
}

func (r *RotatingFileLogger) getFileNameWithoutExt(file string) string {
	i := strings.LastIndex(file, ".")
	return file[:i]
}

func (r *RotatingFileLogger) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (r *RotatingFileLogger) getLogFileList(dir string) ([]logFile, error) {
	var pattern = regexp.MustCompile("_([0-9]*)")
	var files = make([]logFile, 0)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, e := range entries {
		var f logFile
		var filename = r.getFileNameWithoutExt(e.Name())

		matches := pattern.FindAllStringSubmatch(filename, -1)
		if matches != nil {
			logInt, err := strconv.Atoi(matches[0][1])
			if err != nil {
				continue
			}

			f.Filename = filename
			f.LogNumber = logInt
			files = append(files, f)
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].LogNumber < files[j].LogNumber
	})

	return files, nil
}

func (r *RotatingFileLogger) rotateLogs(dir string, filename string, extension string, maxFiles int) error {
	if extension[0:1] != "." {
		extension = "." + extension
	}

	mainLogFile := path.Join(dir, filename+extension)

	fileList, err := r.getLogFileList(dir)
	if err != nil {
		return err
	}

	if r.fileExists(mainLogFile) {
		nextLogNo := 1
		if len(fileList) > 0 {
			nextLogNo = fileList[len(fileList)-1].LogNumber + 1
		}
		nextLogFile := fmt.Sprintf("%s_%d", filename, nextLogNo)
		newLogFile := path.Join(dir, nextLogFile+extension)

		err = os.Rename(mainLogFile, newLogFile)
		if err != nil {
			return err
		}
		fileList = append(fileList, logFile{Filename: nextLogFile, LogNumber: nextLogNo})
	}

	stopRemovingAtIndex := len(fileList) - maxFiles
	newIndex := 1
	for i := 0; i < len(fileList); i++ {
		if i < stopRemovingAtIndex {
			_ = os.Remove(path.Join(dir, fileList[i].Filename+extension))
		} else {
			newFilename := fmt.Sprintf("%s_%d", filename, newIndex)
			_ = os.Rename(
				path.Join(dir, fileList[i].Filename+extension),
				path.Join(dir, newFilename+extension),
			)
			newIndex++
		}
	}

	return nil
}

func (r *RotatingFileLogger) rotate() {
	r.fileLogger.Close()
	err := r.rotateLogs(r.filePath, r.fileName, r.fileExt, r.maxFileCount)
	if err != nil {
		Error(err.Error())
		return
	}
	r.Open()
}

func (r *RotatingFileLogger) Open() {
	logFilePath := path.Join(r.filePath, r.fileName+r.fileExt)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	r.lineCount = countLines(f)
	r.fileLogger.Open(f)
}

func (r *RotatingFileLogger) Close() {
	r.fileLogger.Close()
}

func (r *RotatingFileLogger) SetLevel(level int) {
	r.level = level
	r.fileLogger.level = level
}

func (r *RotatingFileLogger) Log(msg *Message) {
	if msg.level >= r.level {
		r.fileLogger.Log(msg)
		r.lineCount++
		if r.lineCount >= r.maxLines {
			r.rotate()
		}
	}
}

func countLines(f *os.File) int64 {
	r := bufio.NewReader(f)
	var count int64 = 0
	var err error = nil
	for err == nil {
		prefix := true
		_, prefix, err = r.ReadLine()
		if err != nil {
			_ = err
		}
		// sometimes we don't get the whole line at once
		if !prefix && err == nil {
			count++
		}
	}
	return count
}
