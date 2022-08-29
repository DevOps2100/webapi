package global

import (
	//"go.uber.org/zap"
	"webapi/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局参数
var (
	Config config.ServerConfig
	Lg     *zap.Logger
	Trans  ut.Translator
	DB     *gorm.DB
	Redis  *redis.Client
)
