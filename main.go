package main

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"webapi/global"
	"webapi/initialize"
)

// 基于restful风格进行构造
func main() {
	fmt.Println("webapi技巧与方法")
	// 配置初始化
	initialize.InitConfig()
	// MYSQL数据库初始化
	if err := initialize.InitMysqlDB(); err != nil {
		color.Red("数据库初始化异常")
		panic(err)
	}

	// REDIS数据库初始化
	if err := initialize.InitRedis(); err != nil {
		color.Red("数据库初始化异常")
		panic(err)
	}

	// 日志配置初始化
	initialize.InitLogger()

	// websocket启动
	//go func() {
	//	Handler := middlewares.SshHandler{}
	//	http.HandleFunc("/ws", Handler.WebSocket)
	//	http.ListenAndServe(":8080", nil)
	//}()

	// 路由配置初始化
	Router := initialize.Routers()
	color.Green("启动成功")
	// 服务启动
	err := Router.Run(fmt.Sprintf(":%d", global.Config.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误"))
	}

}
