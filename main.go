package main

import (
	"fmt"
	"graphql-parser/controller"

	_ "graphql-parser/docs" // This is required for the generated docs to be included

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Define routes
	r.POST("/parse-graphql", controller.ParseGraphQLSchema)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	r.Run(port)
}
