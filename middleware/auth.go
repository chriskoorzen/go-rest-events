package middleware

import (
	"net/http"

	"github.com/chriskoorzen/go-rest-events/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateJWT(context *gin.Context) {
	// Add authorisation check using JWT
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}
	userID, err := utils.VerifyJWToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}
	context.Set("userID", userID)

	context.Next() // Process request
}
