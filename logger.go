package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ILogger interface {
	Debug(messages string)
	Debugf(message string, value ...interface{})
	Info(message string)
	Infof(message string, value ...interface{})
}

type Logger struct {
	Filepath   string
	FileHandle *os.File
}

func New(logLocation string) ILogger {
	var filePath string
	var fileHandle *os.File

	if logLocation == "" {
		logLocation = filepath.Join("./log", logLocation)
	}

	logFile, err := os.OpenFile(logLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	filePath = logFile.Name()
	fileHandle = logFile
	return &Logger{
		Filepath:   filePath,
		FileHandle: fileHandle,
	}
}

// Info logs things that users care about when using your software.
func (logger *Logger) Info(message string) {
	log(logger, message)
}

// Info logs things that users care about when using your software.
func (logger *Logger) Infof(message string, value ...interface{}) {
	log(logger, fmt.Sprintf(message, value...))
}

// Debug logs things that developers care about when they are developing or debugging software.
func (logger *Logger) Debug(message string) {
	log(logger, message)
}

// Debug logs things that developers care about when they are developing or debugging software.
func (logger *Logger) Debugf(message string, value ...interface{}) {
	log(logger, fmt.Sprintf(message, value...))
}

func log(logger *Logger, message string) {
	logMessageData := time.Now().Format("2006/01/02 15:04:05 ") + message + "\n"
	fmt.Println(logMessageData)
	_, err := logger.FileHandle.WriteString(logMessageData)
	if err != nil {
		panic(err)
	}
}
