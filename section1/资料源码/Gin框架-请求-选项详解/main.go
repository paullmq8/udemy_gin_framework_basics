package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	//router.GET("/someGet",getting)
	//router.POST("/somePost",posting)

	router.Run()
}
