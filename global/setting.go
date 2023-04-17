package global

import (
	"go_api_framework/pkg/logger"
	"go_api_framework/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DataBaseSetting

	Logger *logger.Logger
)
