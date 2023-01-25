package main

import (
	_ "hello/docs"
	"hello/router"
)

// @title Gin swagger
// @version 1.0
// @description Gin Swagger 项目

// @contact.name kangshaojun
// @contact.url www.kangshaojun.com
// @contact.email kangshaojun@gmail.com

// @host localhost:8080
func main() {

	router := router.InitRouter()

	router.Run(":8080")

}
