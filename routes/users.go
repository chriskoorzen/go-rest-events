package routes

import (
	"net/http"

	"github.com/chriskoorzen/go-rest-demo/models"
	"github.com/chriskoorzen/go-rest-demo/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	devOutputBodyToConsole(context) // output the raw body to console

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"error":   err.Error(),
		})
		return
	}

	// If binding is successful, try to save the user
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
			"error":   err.Error(),
		})
		return
	}

	// return newly created user
	context.JSON(http.StatusCreated, gin.H{
		"message": "New user signed up",
		"user":    user,
	})
}

func loginUser(context *gin.Context) {
	devOutputBodyToConsole(context) // output the raw body to console

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"error":   err.Error(),
		})
		return
	}

	// If binding is successful, try to validate the user
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not login user",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
			"error":   err.Error(),
		})
		return
	}

	// return success message
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
