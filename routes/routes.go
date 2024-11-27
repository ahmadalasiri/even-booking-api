package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/auth/signup", signup)
	router.POST("/auth/login", login)

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
}
