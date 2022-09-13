package initialize

import (
	"webapi/middlewares"
	"webapi/router"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 日志和错误处理
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 跨域问题解决
	Router.Use(middlewares.Cors())

	// 服务默认监测路由
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "服务运行正常",
		})
	})
	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.AuthRouter(ApiGroup)
	router.DataRouter(ApiGroup)
	color.Green("路由初始化成功")
	return Router
}
