package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("授权中间件...")
		c.Next()
	}
}
