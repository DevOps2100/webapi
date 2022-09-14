package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webapi/api/forms"
)

var user *forms.Host = nil

func TestDD(ctx *gin.Context) {
	fmt.Println(user, &forms.Host{Name: "123123"})
	fmt.Printf("%T %T\n", user, forms.Host{})

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
