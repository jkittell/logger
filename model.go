package main

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	LogEntry LogEntry
}

type Log struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogEntry struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
