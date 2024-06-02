package global

import (
	"blog/pkg"
	"blog/pkg/logger"
)

var (
	// 配置设置
	ServerSetting   *pkg.ServerSettings
	AppSetting      *pkg.AppSettings
	DatabaseSetting *pkg.DatabaseSettings
	// 日志设置
	Logger *logger.Logger
)
