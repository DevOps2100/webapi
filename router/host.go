package router

import (
	"webapi/api/controller"
	"webapi/middlewares"

	"github.com/gin-gonic/gin"
)

func HostRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	User := Router.Group("host")
	{
		User.POST("add", middlewares.JWTAuth(), controller.AddHost)
		User.POST("get", middlewares.JWTAuth(), controller.GetHost)
		User.GET("getall", middlewares.JWTAuth(), controller.GetHostAll)
		User.POST("delete", middlewares.JWTAuth(), controller.DelHost)
		User.POST("update", middlewares.JWTAuth(), controller.UpdateHost)
	}
}
