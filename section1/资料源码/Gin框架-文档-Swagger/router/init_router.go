package router

import (
	"hello/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	userRouter := router.Group("")
	{
		userRouter.GET("/user/:id", handler.GetOne)
		userRouter.GET("/users", handler.GetAll)
		userRouter.POST("/user", handler.Insert)
		userRouter.DELETE("/user/:id", handler.DeleteOne)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
