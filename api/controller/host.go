package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webapi/api/dao"
	"webapi/api/forms"
	"webapi/api/models"
)

// @Summary 添加主机
// @Tags 主机管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Host body forms.Host true "主机信息"
// @Success 200
// @Router /v1/host/add [post]
// 主机添加
func AddHost(ctx *gin.Context) {
	var host models.Host
	if err := ctx.BindJSON(&host); err != nil {
		fmt.Println("数据绑定失败")
		ctx.JSON(400, gin.H{
			"msg": "数据绑定失败",
		})
		return
	}
	ok, response := dao.AddHost(host)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "添加成功",
	})
}

// @Summary 删除主机
// @Tags 主机管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Host body forms.Host true "主机信息"
// @Success 200
// @Router /v1/host/del [post]
// 主机添加
func DelHost(ctx *gin.Context) {
	var host forms.Host
	if err := ctx.BindJSON(&host); err != nil {
		fmt.Println("数据绑定失败")
		ctx.JSON(400, gin.H{
			"msg": "数据绑定失败",
		})
		return
	}
	ok, response := dao.DelHost(host.Name)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "删除成功",
	})
}

// @Summary 更新主机
// @Tags 主机管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Host body forms.Host true "主机信息"
// @Success 200
// @Router /v1/host/update [post]
// 更新主机
func UpdateHost(ctx *gin.Context) {
	var host forms.Host
	if err := ctx.BindJSON(&host); err != nil {
		fmt.Println("数据绑定失败")
		ctx.JSON(400, gin.H{
			"msg": "数据绑定失败",
		})
		return
	}
	ok, response := dao.UpdateHost(host)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "添加成功",
	})
}

// @Summary 获取主机(单个)
// @Tags 主机管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Host body forms.Host true "主机信息"
// @Success 200
// @Router /v1/host/get [post]
func GetHost(ctx *gin.Context) {
	var host forms.Host
	if err := ctx.BindJSON(&host); err != nil {
		fmt.Println("数据绑定失败")
		ctx.JSON(400, gin.H{
			"msg": "数据绑定失败",
		})
		return
	}
	data, response := dao.GetHost(host.Name)
	if data == nil {
		ctx.JSON(400, gin.H{
			"error": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// @Summary 获取主机(全部)
// @Tags 主机管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Host body forms.Host true "主机信息"
// @Success 200
// @Router /v1/host/getall [get]
func GetHostAll(ctx *gin.Context) {
	var host forms.Host
	if err := ctx.BindJSON(&host); err != nil {
		fmt.Println("数据绑定失败")
		ctx.JSON(400, gin.H{
			"msg": "数据绑定失败",
		})
		return
	}
	data := dao.GetHostAll()
	if data == nil {
		ctx.JSON(400, gin.H{
			"error": "数据为空",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": data,
	})
}
