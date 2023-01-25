package main

import "hello/router"

func main() {

	router := router.InitRouter()

	router.Run(":8080")

}
