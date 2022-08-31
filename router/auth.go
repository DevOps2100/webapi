package router

import (
	"webapi/api/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/auth/login
	Auth := Router.Group("auth")
	{
		Auth.POST("login", controller.Login)
	}
}
