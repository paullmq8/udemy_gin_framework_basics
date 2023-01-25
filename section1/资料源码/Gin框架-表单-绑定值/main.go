package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func main() {

	router := gin.Default()

	router.Static("/", "./public")

	router.POST("/login", func(c *gin.Context) {
		var user UserModel

		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("绑定值发生错误:%s", err.Error()))
		}

		c.String(http.StatusOK, fmt.Sprintf("绑定值成功 %s %s", user.Name, user.Email))
	})

	router.Run(":8080")

}
