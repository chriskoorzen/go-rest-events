package routes

import (
	"net/http"

	"github.com/chriskoorzen/go-rest-demo/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

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

	server.POST("/users", createUser)
	server.POST("/users/login", loginUser)
}
