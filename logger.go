package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	filepath   string
	fileHandle *os.File
}

func New(dir, filename string) *Logger {
	var fileHandle *os.File

	location := filepath.Join(dir, filename)

	logFile, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileHandle = logFile
	return &Logger{
		filepath:   location,
		fileHandle: fileHandle,
	}
}

// Info logs things that users care about when using your software.
func (logger *Logger) Info(value ...any) {
	log(logger, fmt.Sprintln(value))
}

// Infof logs things that users care about when using your software.
func (logger *Logger) Infof(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(logger, msg)
}

// Debug logs things that developers care about when they are developing or debugging software.
func (logger *Logger) Debug(value ...any) {
	log(logger, fmt.Sprintln(value))
}

// Debugf logs things that developers care about when they are developing or debugging software.
func (logger *Logger) Debugf(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(logger, msg)
}

func log(logger *Logger, message string) {
	logMessageData := time.Now().Format("2006/01/02 15:04:05 ") + message
	_, err := logger.fileHandle.WriteString(logMessageData)
	if err != nil {
		panic(err)
	}
}
