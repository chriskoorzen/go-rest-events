package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/chriskoorzen/go-rest-demo/db"
	"github.com/chriskoorzen/go-rest-demo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, REST API!")

	// Init Database
	db.InitDB()

	// Init Server
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
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	devOutputBodyToConsole(context) // output the raw body to console

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"error":   err.Error(),
		})
		return
	}

	// If binding is successful, try to save the event
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event",
			"error":   err.Error(),
		})
		return
	}

	// return newly created event
	context.JSON(http.StatusCreated, gin.H{
		"message": "POST successful",
		"event":   event,
	})
}

func devOutputBodyToConsole(context *gin.Context) {
	// output the raw body for dev purposes
	body, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the request body
	fmt.Println("Raw Body:", string(body))
}
