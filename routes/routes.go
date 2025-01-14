package routes

import (
	"net/http"

	"github.com/chriskoorzen/go-rest-events/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(middleware.ErrorHandler)

	protected := server.Group("/")
	protected.Use(middleware.AuthenticateJWT)

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello REST API"})
	})

	server.GET("/events", getEvents)
	protected.POST("/events", createEvent)
	server.GET("/events/:eventsID", getSingleEvent)
	protected.PUT("/events/:eventsID", updateEvent)
	protected.DELETE("/events/:eventsID", deleteEvent)

	protected.GET("/events/:eventsID/register", getEventRegistrations)
	protected.POST("/events/:eventsID/register", registerForEvent)
	protected.DELETE("/events/:eventsID/register", cancelRegistration)

	server.POST("/users", createUser)
	server.POST("/users/login", loginUser)
}
