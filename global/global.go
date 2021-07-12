package global

import (
	"gin_admin_api/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// ServerConfig 配置项
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	// Logger 日志
	Logger *zap.Logger
	// DB 数据库
	DB *gorm.DB
	// Trans 数据校验
	//Trans ut.Translator
)
