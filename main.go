package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, REST API!")
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello REST API",
		})
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET all events",
	})
}

func createEvent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POSTed an event",
	})
}
