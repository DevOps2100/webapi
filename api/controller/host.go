package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webapi/api/dao"
	"webapi/api/forms"
)

// 主机添加
func AddHost(ctx *gin.Context) {
	var host forms.Host
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
