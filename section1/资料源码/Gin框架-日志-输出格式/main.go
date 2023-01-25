package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("自定义日志:%s - [%s] %s %s %s %s", param.ClientIP, param.TimeStamp, param.Method, param.Request.Proto, param.StatusCode)
	}))

	router.GET("/testLog", func(c *gin.Context) {
		c.String(200, "hello")
	})

	router.Run(":8080")

}
