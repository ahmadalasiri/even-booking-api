package main

import (
	"event-booking-api/db"
	"event-booking-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(
		":4001",
	)
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(200, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = len(models.GetAllEvents()) + 1
	event.DateTime = time.Now()
	event.UserID = int(uuid.New().ID())

	event.Save()
	c.JSON(200, event)
}
