package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/testJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})

	r.GET("/testJSON1", func(c *gin.Context) {
		var message struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		message.Name = "kevin"
		message.Message = "hello"
		message.Number = 123456

		c.JSON(http.StatusOK, message)

	})

	r.GET("/testXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})

	r.GET("/testYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})

	r.GET("/testHTML", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello world</b>",
		})
	})

	r.Run(":8080")

}
