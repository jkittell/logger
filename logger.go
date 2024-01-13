package main

import (
	"context"
	"github.com/jkittell/data/database"
	"log"
	"time"
)

type LogEntry struct {
	Name    string `json:"name"`
	Message string `bson:"message" json:"message"`
}

type logEntry struct {
	Message   string    `bson:"message" json:"message"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type logger struct {
	logs database.MongoDB[logEntry]
}

func newLogger() (*logger, error) {
	logs, err := database.NewMongoDB[logEntry]()
	return &logger{logs: logs}, err
}

func (l logger) log(entry LogEntry) {
	err := l.logs.Insert(context.TODO(), entry.Name, logEntry{
		Message:   entry.Message,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
}
