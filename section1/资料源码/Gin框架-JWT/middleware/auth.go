package middleware

import (
	"hello/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {
		result := model.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证,请重新登录",
			Data:    nil,
		}

		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}

		auth = strings.Fields(auth)[1]
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Message = "token过期" + err.Error()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		} else {
			println("token验证通过")
		}
		context.Next()
	}

}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		//密码
		return []byte("gin hello"), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
