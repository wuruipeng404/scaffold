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

func log() *zap.SugaredLogger {
	if _log == nil {
		_log = NewSugarLogger(util.IF(os.Getenv("SCAFFOLD_LOG_PATH"), util.DefaultLogPath))
	}
	return _log
}

// DefaultLogger default logger with name
func DefaultLogger(path string) *zap.SugaredLogger {
	_log = NewSugarLogger(path)
	return _log
}

func Debug(args ...interface{}) {
	log().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log().Debugf(template, args...)
}

func Info(args ...interface{}) {
	log().Info(args...)
}

func Infof(template string, args ...interface{}) {
	log().Infof(template, args...)
}

func Error(args ...interface{}) {
	log().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log().Errorf(template, args...)
}

func Warn(args ...interface{}) {
	log().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	_log.Warnf(template, args...)
}

func Panic(args ...interface{}) {
	log().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log().Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	log().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	log().Fatalf(template, args...)
}

func Sync() {
	_ = log().Sync()
}
