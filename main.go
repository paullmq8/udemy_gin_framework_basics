package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `format:"user" binding:"required"`
	Password string `format:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"status": "ok"})))
	router.POST("/login", test)
	router.Run(":8080")
}

func test(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
	c.JSON(http.StatusMovedPermanently, gin.H{"status": "moved permanently"})
	c.Redirect()
	c.Request.URL.Path = "/login"
}
