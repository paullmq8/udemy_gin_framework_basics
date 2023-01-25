package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.GET("/testLog", func(c *gin.Context) {
		c.String(200, "hello")
	})

	router.Run(":8080")

}
