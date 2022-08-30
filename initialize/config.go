package initialize

import (
	"webapi/config"
	"webapi/global"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 使用viper来对配置进行实例化
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		color.Red("读取配置文件失败， 请检查配置是否正确！")
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		color.Red("格式化文件失败， 请检查配置字段对接是否正确！")
		panic(err)
	}

	global.Config = serverConfig
	color.Green("配置初始化成功")
}
