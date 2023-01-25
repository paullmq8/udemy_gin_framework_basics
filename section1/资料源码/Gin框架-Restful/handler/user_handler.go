package handler

import (
	"hello/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	user := model.UserModel{}
	users := user.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id不是int类型", e.Error())
	}

	user := model.UserModel{
		Id: i,
	}

	u := user.FindById()
	context.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func Insert(context *gin.Context) {
	user := model.UserModel{}

	var id = -1
	if e := context.ShouldBindJSON(&user); e == nil {
		id = user.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id不是int类型", e.Error())
	}

	user := model.UserModel{
		Id: i,
	}
	user.DeleteOne()
}
