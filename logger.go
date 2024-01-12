package main

import (
	"context"
	"github.com/jkittell/data/database"
	"log"
	"time"
)

type LogEntry struct {
	Log       any       `bson:"log" json:"log"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type Logger struct {
	logs database.MongoDB[LogEntry]
}

func NewLogger() (*Logger, error) {
	logs, err := database.NewMongoDB[LogEntry]("logs")
	return &Logger{logs: logs}, err
}

func (l Logger) WriteLog(data any) {
	err := l.logs.Insert(context.TODO(), LogEntry{
		Log:       data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
}
