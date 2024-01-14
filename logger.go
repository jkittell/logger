package main

import (
	"context"
	"github.com/jkittell/data/database"
	"log"
	"os"
	"time"
)

type logEntry struct {
	Message   string    `bson:"message" json:"message"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}

type logger struct {
	collectionName string
	logs           database.MongoDB[logEntry]
}

func newLogger() (*logger, error) {
	logs, err := database.NewMongoDB[logEntry]()
	return &logger{collectionName: os.Getenv("MONGODB_COLLECTION_NAME"), logs: logs}, err
}

func (l logger) log(message string) {
	err := l.logs.Insert(context.TODO(), l.collectionName, logEntry{
		Message:   message,
		Timestamp: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
}
