package main

import (
	"context"
	"github.com/jkittell/data/database"
	"log"
	"time"
)

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

func (l logger) log(message string) {
	err := l.logs.Insert(context.TODO(), "logs", logEntry{
		Message:   message,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
}
