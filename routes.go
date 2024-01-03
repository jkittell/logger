package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jkittell/data/database"
	"net/http"
	"time"
)

type LogHandler struct {
	logs database.MongoDB[LogEntry]
}

func NewLogHandler() (*LogHandler, error) {
	logs, err := database.NewMongoDB[LogEntry](".env", "logs")
	return &LogHandler{logs: logs}, err
}

func (h LogHandler) WriteLog(c *gin.Context) {
	var log Log

	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := LogEntry{
		ID:        uuid.New(),
		Name:      log.Name,
		Data:      log.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := h.logs.Insert(context.TODO(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error inserting log"})
		return
	}

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"error:": false, "message": "logged"})
}
