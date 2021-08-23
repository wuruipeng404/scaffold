/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 10:23
 */

package logger

import (
	"os"

	"github.com/wuruipeng404/scaffold/util"
	"go.uber.org/zap"
)

var _log *zap.SugaredLogger

// 默认生成一个日志对象
func init() {
	_log = NewSugarLogger(util.IF(os.Getenv("SCAFFOLD_LOG_PATH"), "log/default.log"))
}

func DefaultLogger() *zap.SugaredLogger {
	return _log
}

func Debug(args ...interface{}) {
	_log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	_log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	_log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	_log.Infof(template, args...)
}

func Error(args ...interface{}) {
	_log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	_log.Errorf(template, args...)
}

func Warn(args ...interface{}) {
	_log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	_log.Warnf(template, args...)
}

func Panic(args ...interface{}) {
	_log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	_log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	_log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	_log.Fatalf(template, args...)
}

func Sync() {
	_ = _log.Sync()
}
