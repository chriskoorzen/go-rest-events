package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello REST API"})
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:eventsID", getSingleEvent)
	server.PUT("/events/:eventsID", updateEvent)
	server.DELETE("/events/:eventsID", deleteEvent)

	server.POST("/users", createUser)
	server.POST("/users/login", loginUser)
}
