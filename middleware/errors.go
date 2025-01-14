package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(context *gin.Context) {
	context.Next() // Process request

	// Collect and output all errors
	if len(context.Errors) > 0 {
		errors := make([]string, len(context.Errors))
		for i, err := range context.Errors {
			errors[i] = err.Error()
		}

		// Just output to console for now
		fmt.Println("[Errors]: ", errors)
	}
}
