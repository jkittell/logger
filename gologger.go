package gologger

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

func NewLogger(dir, filename string) *Logger {
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
func (Logger *Logger) Info(value ...any) {
	log(Logger, fmt.Sprintln(value))
}

// Info logs things that users care about when using your software.
func (Logger *Logger) Infof(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(Logger, msg)
}

// Debug logs things that developers care about when they are developing or debugging software.
func (Logger *Logger) Debug(value ...any) {
	log(Logger, fmt.Sprintln(value))
}

// Debug logs things that developers care about when they are developing or debugging software.
func (Logger *Logger) Debugf(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(Logger, msg)
}

func log(Logger *Logger, message string) {
	logMessageData := time.Now().Format("2006/01/02 15:04:05 ") + message
	fmt.Println(logMessageData)
	_, err := Logger.fileHandle.WriteString(logMessageData)
	if err != nil {
		panic(err)
	}
}
