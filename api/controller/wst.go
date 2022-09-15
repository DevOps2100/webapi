package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webapi/api/forms"
)

func WEBSET(ctx *gin.Context) {
	info := forms.WSHost{
		Username:  "root",
		Password:  "",
		Port:      22,
		Ipaddress: "42.193.154.221",
	}
	fmt.Println(info)
	ctx.HTML(200, "/static/front/inedex.html", gin.H{})
}
