package router

import (
	"hello/handler"
	"hello/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/login", handler.LoginHandler)

	router.GET("/testToken", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})

	return router

}
