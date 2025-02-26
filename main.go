package main

import (
	"fmt"

	"github.com/chriskoorzen/go-rest-events/db"
	"github.com/chriskoorzen/go-rest-events/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, REST API!")

	// Init Database
	db.InitDB()

	// Init Server
	server := gin.Default()

	// Init Routes
	routes.RegisterRoutes(server)

	// Start Server
	server.Run(":8080")
}
