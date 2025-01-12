package routes

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func devOutputBodyToConsole(context *gin.Context) {
	// output the raw body for dev purposes
	body, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the request body
	fmt.Println("Raw Body:", string(body))
}
