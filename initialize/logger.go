package initialize

import (
	"fmt"
	"time"

	"github.com/fatih/color"

	"webapi/global"

	"go.uber.org/zap"
)

// InitLogger 初始化Logger
func InitLogger() {
	// 实例化zap配置
	cfg := zap.NewDevelopmentConfig()

	// 配置日志输出的地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog%s.log", global.Config.LogsAddress, GetNowFormatTodayTime()),
		"stdout",
	}

	// 创建logger实例
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
	color.Green("日志配置初始化成功")
}

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}
