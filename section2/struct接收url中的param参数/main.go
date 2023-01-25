package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

// call this function by:
// curl -X GET "http://localhost:8888/user?name=paul&address=canada" -H "Content-Type: application/json" -d '{"name": "paul", "address": "canada"}'
// or visit http://localhost:8888/user?name=paul&address=canada in browser and check the difference
func main() {
	router := gin.Default()
	router.GET("/user", test)
	router.Run(":8888")
}

func test(c *gin.Context) {
    var user User
    if c.ShouldBindQuery(&user) == nil {
        log.Println(user.Name, ",", user.Address)
        c.JSON(http.StatusOK, gin.H{"status1": "ok", "name": user.Name, "address": user.Address})
    }
	if c.ShouldBindJSON(&user) == nil {
		log.Println(user.Name, ",", user.Address)
		c.JSON(http.StatusOK, gin.H{"status2": "ok", "name": user.Name, "address": user.Address})
	}
}
