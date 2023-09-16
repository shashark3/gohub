package logger

import (
	"context"
	"errors"
	"gohub/pkg/helpers"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// GormLogger 操作对象，实现gormlogger.interface
type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

// NewGormLogger 外部调用。实例化一个 GormLogger 对象，示例：
//
//	DB, err := gorm.Open(dbConfig, &gorm.Config{
//	    Logger: logger.NewGormLogger(),
//	})
func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,                 //使用全局的Logger变量
		SlowThreshold: 200 * time.Millisecond, //慢查询阈值
	}
}

// LogMode 实现gormlogger.Interface的LogMode方法
func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

// Info 实现 gormlogger.Interface 的 Info 方法
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)
}

// Warn 实现 gormlogger.Interface 的 Warn 方法
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}

// Error 实现 gormlogger.Interface 的 Error 方法
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args...)
}

// Trace 实现 gormlogger.Interface 的 Trace 方法
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	//获取运行时间
	elapsed := time.Since(begin)
	//获取Sql请求和返回条数
	sql, rows := fc()

	//通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", helpers.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}
	//Gorm错误
	if err != nil {
		//记录未找到的错误使用warning等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)

		} else {
			//其他错误使用error等级
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database Error", logFields...)

		}

	}
	//慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow log", logFields...)
	}
	//	记录所有Sql请求
	l.logger().Debug("Database Query", logFields...)

}

// logger 内用的辅助方法，确保 Zap 内置信息 Caller 的准确性（如 paginator/paginator.go:148）
func (l GormLogger) logger() *zap.Logger {
	//跳过Gorm 内置的使用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	//减去一次封装，以及一次在Logger初始化里添加zap.AddCallerskip(1)
	clone := l.ZapLogger

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return l.ZapLogger
}
