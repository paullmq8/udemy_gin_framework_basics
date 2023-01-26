package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Call this function with the command: go run main.go
// curl -X POST "http://localhost:8888/login" -H "Content-Type: application/json" -d '{"user": "user", "password": "password1"}'
// curl -v --form user=user --form password=password1 http://localhost:8888/login
// These two curls are different for c.PostForm
func main() {
	router := gin.Default()
	router.POST("/login", test)
	router.Run(":8888")
}

func test(c *gin.Context) {
    u := c.PostForm("user")
    fmt.Printf("user: %s\n", u)

    var form LoginForm
    if c.ShouldBind(&form) == nil {
        if form.User == "user" && form.Password == "password" {
            c.JSON(http.StatusOK, gin.H{"status": "Success"})
        } else {
            c.JSON(http.StatusForbidden, gin.H{"status": "Error"})
        }
    }
}