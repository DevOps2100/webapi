package router

import (
	"webapi/api/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	User := Router.Group("user")
	{
		User.GET("login", controller.Login)
		User.POST("login", controller.Login)
	}
}
