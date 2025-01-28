package routes

import (
	"net/http"

	"github.com/chriskoorzen/go-rest-events/models"
	"github.com/chriskoorzen/go-rest-events/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// If binding is successful, try to save the user
	err = user.Save()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// return newly created user
	context.JSON(http.StatusCreated, gin.H{
		"message": "New user successfully created",
	})
}

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse POST request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// If binding is successful, try to validate the user
	err = user.ValidateCredentials()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not login user",
			"code":    http.StatusUnauthorized,
		})
		return
	}

	token, err := utils.GenerateJWToken(user.ID, user.Email)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// return success message
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
