package controller

import (
	"fmt"
	"webapi/api/dao"
	"webapi/api/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户添加
func UserAdd(ctx *gin.Context) {
	var user forms.UserAddForm
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("数据格式错误")
		ctx.JSON(400, gin.H{
			"msg": "error",
		})
		return
	} else {
		response := dao.AddUser(&user)
		fmt.Println(response)
		zap.L().Info(response)
		ctx.JSON(200, gin.H{
			"msg": response,
		})
	}
}

// 用户查询(单个)
func GetUser(ctx *gin.Context) {
	var user forms.UsernameForm
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("数据获取错误")
		ctx.JSON(400, gin.H{
			"msg": err,
		})
		return
	} else {
		response, str := dao.GetUser(user.Username)
		if str == "用户不存在" {
			ctx.JSON(200, gin.H{
				"msg": str,
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": response,
				"msg":  str,
			})
		}

	}
}

// 用户查询（全部）
func GetUserAll(ctx *gin.Context) {
	response := dao.GetUserAll()
	ctx.JSON(200, gin.H{
		"data": response,
	})
}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	var user forms.UsernameForm
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("数据参数错误")
		ctx.JSON(200, gin.H{
			"msg": "数据参数错误",
		})
		return
	} else {
		ok, response := dao.DeleteUser(user.Username)
		if ok {
			zap.L().Info(response)
			ctx.JSON(200, gin.H{
				"msg": response,
			})
		} else {
			zap.L().Info(response)
			ctx.JSON(200, gin.H{
				"msg": response,
			})
		}
	}
}

// 用户修改
func UpdateUser(ctx *gin.Context) {
	var user forms.UserInfo
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("用户数据错误")
		ctx.JSON(200, gin.H{
			"msg": "用户数据错误",
		})
		return
	}
	ok, response := dao.UpdateUser(user.Username, user.Password)
	if ok {
		zap.L().Info("用户修改成功")
		ctx.JSON(200, gin.H{
			"msg": response,
		})
		return
	}
}
