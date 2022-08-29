package initialize

import (
	"fmt"
	"webapi/global"

	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func InitRedis() error {
	addr := fmt.Sprintf("%s:%d", global.Config.RedisInfo.Host, global.Config.RedisInfo.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.Config.RedisInfo.Password,
		DB:       global.Config.RedisInfo.DB,
	})

	// 链接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
		return err
	}
	color.Green("REDIS数据库初始化成功")
	return nil
}
