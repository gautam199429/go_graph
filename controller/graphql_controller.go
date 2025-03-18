package controller

import (
	"graphql-parser/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ParseGraphQLSchema godoc
// @Summary Parse GraphQL Schema
// @Description Parse a given GraphQL schema and return the types
// @Tags schema
// @Accept  json
// @Produce  json
// @Param schema body string true "GraphQL Schema"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /parse-graphql [post]
func ParseGraphQLSchema(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": 0, "message": "Error reading body"})
		return
	}
	typeMap, err := service.ParseSchema(string(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": 0, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 1,
		"message":     "Query parsed successfully",
		"data":        typeMap,
	})
}
