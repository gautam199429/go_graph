package controller

import (
	"graphql-parser/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
