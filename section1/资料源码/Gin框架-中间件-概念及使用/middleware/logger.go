package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		fmt.Printf("%s  %s  %s \n", time.Now().Format("2006-01-01 11:11:11"), host, url)
		context.Next()
	}
}
