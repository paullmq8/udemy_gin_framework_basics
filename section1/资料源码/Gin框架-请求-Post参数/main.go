package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginFrom struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", test)
	router.Run(":8080")
}

func test(c *gin.Context) {

	u := c.PostForm("user")
	fmt.Printf("用户的值是:%s\n", u)

	var form LoginFrom
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"状态": "登录成功"})
		} else {
			c.JSON(200, gin.H{"状态": "登录失败"})
		}
	}

}

//curl -v --form user=user --form password=password http://localhost:8080/login
