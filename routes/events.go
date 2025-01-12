package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/chriskoorzen/go-rest-demo/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get events",
			"error":   err.Error(),
		})
		return
	}

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

func getSingleEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func devOutputBodyToConsole(context *gin.Context) {
	// output the raw body for dev purposes
	body, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the request body
	fmt.Println("Raw Body:", string(body))
}
