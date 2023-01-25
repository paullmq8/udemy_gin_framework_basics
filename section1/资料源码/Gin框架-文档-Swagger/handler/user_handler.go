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

// @Summary 通过Id查询用户
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 {object} model.UserModel 成功后返回值
// @Router /user/{id} [get]
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

// @Summary 添加用户
// @Id 1
// @Tags 用户
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.UserModel true "用户"
// @Success 200 object model.UserModel 成功后返回值
// @Failure 409 object model.UserModel 添加失败
// @Router /user [post]
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

// @Summary 通过id删除用户
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200  object model.UserModel 成功后返回值
// @Router /user/{id} [delete]
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
