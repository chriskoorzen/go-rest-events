package routes

import (
	"net/http"
	"strconv"

	"github.com/chriskoorzen/go-rest-events/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get events",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Retrieve userID from context
	userID := context.GetInt64("userID")
	event.CreatorID = userID // connect event to specific user

	// If binding is successful, try to save the event
	err = event.Save()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event",
			"code":    http.StatusInternalServerError,
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
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
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
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// check if user is authorised to update event -> is the creator of the event
	if event.CreatorID != context.GetInt64("userID") {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event",
			"code":    http.StatusUnauthorized,
		})
		return
	}

	// attempt to update event
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated event"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// check if user is authorised to delete event -> is the creator of the event
	if event.CreatorID != context.GetInt64("userID") {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event",
			"code":    http.StatusUnauthorized,
		})
		return
	}

	// attempt to delete event
	err = event.Delete()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted event"})
}

func registerForEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// attempt to register for event
	userID := context.GetInt64("userID")
	err = event.Register(userID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register for event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func cancelRegistration(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// attempt to cancel registration
	userID := context.GetInt64("userID")
	err = event.CancelRegistration(userID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration for event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration for event"})
}

func getEventRegistrations(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventsID"), 10, 64)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// check if user is authorised to view registrations
	// -> is the creator of the event
	if event.CreatorID != context.GetInt64("userID") {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to view registrations for event",
			"code":    http.StatusUnauthorized,
		})
		return
	}

	registrations, err := event.GetRegistrations()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get registrations for event",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"registrations": registrations,
	})
}
