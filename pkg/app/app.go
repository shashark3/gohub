package app

//package app 应用信息

import (
	"gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimenowInTimeZone 获取当前时间，支持时区
func TimenowInTimeZone() time.Time {
	ChinaTimeZone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(ChinaTimeZone)
}
