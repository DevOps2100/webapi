package router

import (
	"webapi/api/controller"
	"webapi/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	User := Router.Group("user")
	{
		User.POST("add", middlewares.JWTAuth(), controller.UserAdd)
		User.GET("get", middlewares.JWTAuth(), controller.GetUser)
		User.GET("getall", middlewares.JWTAuth(), controller.GetUserAll)
		User.POST("delete", middlewares.JWTAuth(), controller.DeleteUser)
		User.POST("update", middlewares.JWTAuth(), controller.UpdateUser)
	}
}
