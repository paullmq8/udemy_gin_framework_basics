package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Static("/", "./public")

	router.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		c.String(http.StatusOK, fmt.Sprintf("name=%s email=%s", name, email))
	})

	router.Run(":8080")

}
