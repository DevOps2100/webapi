package controller

import (
	"github.com/gin-gonic/gin"
)

// 用户登录
func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "success",
	})
}
