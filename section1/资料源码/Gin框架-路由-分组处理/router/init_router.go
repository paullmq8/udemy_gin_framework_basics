package router

import (
	"hello/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		//http://localhost:8080/user/kevin
		userRouter.GET("/:name", handler.UserSave)
		//http://localhost:8080/user/?name=kevin&age=30
		userRouter.GET("/", handler.UserSaveByQuery)
	}

	return router
}
