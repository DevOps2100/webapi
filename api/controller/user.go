package controller

import (
	"fmt"
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

// @Summary 用户添加
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.UserAddForm body forms.UserAddForm true "用户信息"
// @Success 200
// @Router /v1/user/add [post]
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

// @Summary 用户查询(单个)
// @Tags 用户管理
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param username body forms.UsernameForm true  "用户信息"
// @Success 200
// @Router /v1/user/get [post]
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

// @Summary 用户查询(全部)
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Success 200
// @Router /v1/user/getall [get]
// 用户查询（全部）
func GetUserAll(ctx *gin.Context) {
	response := dao.GetUserAll()
	ctx.JSON(200, gin.H{
		"data": response,
	})
}

// @Summary 用户删除
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.UsernameForm body forms.UsernameForm true "用户信息"
// @Success 200
// @Router /v1/user/delete [post]
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

// @Summary 用户修改
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.UserInfo body forms.UserInfo true "用户信息"
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/user/update [post]
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

	// 密码加密
	password := GeneratePassword(user.Password)
	ok, response := dao.UpdateUser(user.Username, password)
	if ok {
		zap.L().Info("用户修改成功")
		ctx.JSON(200, gin.H{
			"msg": response,
		})
		return
	}
}
