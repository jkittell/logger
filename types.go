package main

import (
	"time"
)

type logEntry struct {
	Message   string    `bson:"message" json:"message"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
