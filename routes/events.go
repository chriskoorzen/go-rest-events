package routes

import (
	"net/http"
	"strconv"

	"github.com/chriskoorzen/go-rest-demo/models"
	"github.com/chriskoorzen/go-rest-demo/utils"
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

	// Add authorisation check
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}
	userID, err := utils.VerifyJWToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}
	// end authorisation check

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"error":   err.Error(),
		})
		return
	}
	event.UserID = userID // connect event to specific user

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

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}

	// check if event exists
	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"error":   err.Error(),
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated event"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"error":   err.Error(),
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted event"})
}
