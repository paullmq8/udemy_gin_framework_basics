package main

import (
	"hello/router"
)

func main() {

	router := router.SetUpRouter()

	router.Run(":8080")

}
