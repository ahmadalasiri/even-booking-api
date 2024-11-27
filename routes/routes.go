package routes

import (
	"event-booking-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/auth/signup", signup)
	router.POST("/auth/login", login)

	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authorization)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)

}
