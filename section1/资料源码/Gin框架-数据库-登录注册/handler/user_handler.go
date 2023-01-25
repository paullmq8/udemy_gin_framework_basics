package handler

import (
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Panicln("注册失败", err.Error())
	}

	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "两次输入密码不一致")
		log.Panicln("两次输入密码不一致")
	}
	id := user.Save()
	log.Panicln("插入的Id为:", id)
	context.Redirect(http.StatusMovedPermanently, "/user/login")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("登录绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
	} else {
		log.Println("登录失败")
	}
}
