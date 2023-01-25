package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "保存用户为:"+username)
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "保存用户为:"+username+" 年龄为:"+age)
}
