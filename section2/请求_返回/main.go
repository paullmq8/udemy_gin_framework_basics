package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.GET("/index", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World!",
        })
        // c.JSON(http.StatusOK, "aaa")
    })

    router.POST("")
    router.Run(":8888")
}