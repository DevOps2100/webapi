package initialize

import (
	"webapi/middlewares"
	"webapi/router"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"webapi/docs"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// swagger 配置
	// programatically set swagger info
	docs.SwaggerInfo.Title = "webapi"
	docs.SwaggerInfo.Description = "go语言基础gin快速开发框架"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// 开启路由，文档访问地址  go1.16版本之后需要安装最新版本： https://github.com/swaggo/swag
	// go install github.com/swaggo/swag/cmd/swag@latest
	Router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

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
	router.HostRouter(ApiGroup)
	color.Green("路由初始化成功")
	return Router
}
