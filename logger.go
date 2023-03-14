package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ILogger interface {
	Debug(value ...any)
	Debugf(message string, value ...any)
	Info(value ...any)
	Infof(message string, value ...any)
}

type logger struct {
	filepath   string
	fileHandle *os.File
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
	return &logger{
		filepath:   filePath,
		fileHandle: fileHandle,
	}
}

// Info logs things that users care about when using your software.
func (logger *logger) Info(value ...any) {
	log(logger, fmt.Sprintln(value))
}

// Info logs things that users care about when using your software.
func (logger *logger) Infof(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(logger, msg)
}

// Debug logs things that developers care about when they are developing or debugging software.
func (logger *logger) Debug(value ...any) {
	log(logger, fmt.Sprintln(value))
}

// Debug logs things that developers care about when they are developing or debugging software.
func (logger *logger) Debugf(message string, value ...interface{}) {
	msg := fmt.Sprintf(message, value...) + "\n"
	log(logger, msg)
}

func log(logger *logger, message string) {
	logMessageData := time.Now().Format("2006/01/02 15:04:05 ") + message
	fmt.Println(logMessageData)
	_, err := logger.fileHandle.WriteString(logMessageData)
	if err != nil {
		panic(err)
	}
}
