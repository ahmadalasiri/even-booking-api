package routes

import (
	"event-booking-api/middlewares"
	"event-booking-api/models"
	"event-booking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)
}

func getEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := models.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, event)
}

func createEvent(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var event models.Event
	event.UserID = user.UserID

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEvent, err := event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEvent)

}

func updateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	userValue, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Type assert to middlewares.User
	user, ok := userValue.(middlewares.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event.UserID != user.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this event"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent, err := event.UpdateEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(c *gin.Context) {
	id := c.Param("id")

	userValue, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Type assert to middlewares.User
	user, ok := userValue.(middlewares.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event.UserID != user.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = models.DeleteEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
