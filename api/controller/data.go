package controller

import (
	"github.com/gin-gonic/gin"
	"webapi/api/dao"
	"webapi/api/forms"
)

// @Summary 添加数据源
// @Tags 数据源管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Data body forms.Data true "数据源信息"
// @Success 200
// @Router /v1/data/add [post]
func DataAdd(ctx *gin.Context) {
	var data forms.Data
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{
			"erros": "数据绑定错误",
		})
		return
	}
	ok, response := dao.DataAdd(data)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": response,
	})
}

// @Summary 删除数据源
// @Tags 数据源管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Data body forms.Data true "数据源信息"
// @Success 200
// @Router /v1/data/delete [post]
func DataDel(ctx *gin.Context) {
	var data forms.Data
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{
			"erros": "数据绑定错误",
		})
		return
	}
	ok, response := dao.DataDel(data.Name)
	if !ok {
		ctx.JSON(400, gin.H{
			"msg": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": response,
	})
}

// @Summary 更新数据源
// @Tags 数据源管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Data body forms.Data true "数据源信息"
// @Success 200
// @Router /v1/data/update [post]
func DataUpdate(ctx *gin.Context) {
	var data forms.Data
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{
			"erros": "数据绑定错误",
		})
		return
	}
	ok, response := dao.DataUpdate(data)
	if !ok {
		ctx.JSON(400, gin.H{
			"msg": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": response,
	})
}

// @Summary 获取数据源
// @Tags 数据源管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Data body forms.Data true "数据源信息"
// @Success 200
// @Router /v1/data/get [post]
func DataGet(ctx *gin.Context) {
	var data forms.Data
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{
			"erros": "数据绑定错误",
		})
		return
	}
	row, response := dao.DataGet(data)
	if row == nil {
		ctx.JSON(400, gin.H{
			"msg": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": row,
	})
}
