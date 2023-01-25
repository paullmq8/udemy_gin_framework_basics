package router

import (
	"hello/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "./statics")

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	return router

}
