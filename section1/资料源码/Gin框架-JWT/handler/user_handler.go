package handler

import (
	"hello/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {

	user := &model.UserModel{}
	result := &model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}

	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {

		expiresTime := time.Now().Unix() + int64(60*60*24)
		claims := jwt.StandardClaims{
			Audience:  user.Email,        // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        string(user.Id),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin hello",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",
		}

		var secret = []byte("gin hello")
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if token, err := tokenClaims.SignedString(secret); err == nil {

			result.Message = "登录成功"
			result.Data = "Bearer " + token
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})

		} else {
			result.Message = "登录失败"
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusOK
		ctx.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}
