package controller

import (
	"github.com/gin-gonic/gin"
	"webapi/api/dao"
	"webapi/api/forms"
)

// 数据源管理

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
	}
	ctx.JSON(200, gin.H{
		"msg": response,
	})
}

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
