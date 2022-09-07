package controller

import (
	"fmt"
	"strings"
	"webapi/api/dao"
	"webapi/api/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

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
		if user.Username == "" || user.Password == "" {
			zap.L().Info("数据格式错误")
			ctx.JSON(400, gin.H{
				"msg": "error",
			})
			return
		}
		result := GeneratePassword(user.Password)
		// fmt.Println("加密密码:", string(result))
		user.Password = string(result)
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
		username := strings.Trim(user.Username, "\"")
		fmt.Println("请求数据： ", user.Username)
		fmt.Println("请求数据： ", username)
		response, str := dao.GetUser(username)
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
	var user forms.UserAddForm
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("用户数据错误")
		ctx.JSON(200, gin.H{
			"msg": "用户数据错误",
		})
		return
	}

	// 密码加密
	user.Password = GeneratePassword(user.Password)
	ok, response := dao.UpdateUser(user)
	if ok {
		zap.L().Info("用户修改成功")
		ctx.JSON(200, gin.H{
			"msg": response,
		})
		return
	}
}
