package setting

import "time"

/*
  在section中定义要使用的配置结构体，通过setting进行解析。
*/

type ServerSetting struct {
	HttpPort     string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	LogTraceEnable  bool
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DataBaseSetting struct {
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
}
