package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/index.html","templates/home.html")

	router.StaticFile("/logo.jpeg", "./logo.jpeg")
	router.Static("/assets", "./assets")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "标题",
			"message": "你好",
		})
	})

	router.Run(":8080")

}
