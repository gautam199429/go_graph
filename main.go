package main

import (
	"fmt"
	"graphql-parser/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes
	r.POST("/parse-graphql", controller.ParseGraphQLSchema)

	// Start the server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	r.Run(port)
}
