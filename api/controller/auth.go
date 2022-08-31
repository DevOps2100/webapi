package controller

import (
	"fmt"
	"webapi/api/dao"
	"webapi/api/forms"
	"webapi/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type GenTokenInfo struct {
	Token     string
	ExpiresAt int64
}

// 用户登录
func Login(ctx *gin.Context) {
	var user forms.LoginForm
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("login error: ", zap.Error(err))
		ctx.JSON(401, gin.H{
			"msg": "格式错误",
		})
		return
	}
	model := dao.GetUserByUsername(user.Username)
	fmt.Println("用户信息: ", model.Username, model.Password)
	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("对比错误： ", err)
		ctx.JSON(401, gin.H{
			"msg": "AuthError",
		})
	} else {
		expires, token, _ := middlewares.GenToken(user.Username)
		data := GenTokenInfo{
			Token:     token,
			ExpiresAt: expires,
		}
		ctx.JSON(200, gin.H{
			"msg":  "登陆成功",
			"data": data,
		})
		return
	}
}
