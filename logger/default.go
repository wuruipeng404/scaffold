/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 10:23
 */

package logger

import (
	"go.uber.org/zap"
)

var _log *zap.SugaredLogger

// InitLogger InitDefaultLogger default logger with name, should be init call it
func InitLogger(path string) {
	_log = newSugarLogger(path, 1)
}

func NewLogger(path string, skip int) *zap.SugaredLogger {
	return newSugarLogger(path, skip)
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
	if _log != nil {
		_ = _log.Sync()
	}
}
