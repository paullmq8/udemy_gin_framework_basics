package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.GET("/user", test)
	router.Run(":8080")
}

func test(c *gin.Context) {
	var user User
	if c.ShouldBindQuery(&user) == nil {
		log.Println(user.Name)
		log.Println(user.Address)
	}
	c.String(200, "成功")
}
