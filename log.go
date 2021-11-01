/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/10/28 14:46
 */

package scaffold

import (
	"os"

	"github.com/wuruipeng404/scaffold/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	log = newConsoleLogger()
}

func newConsoleLogger() *zap.SugaredLogger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     logger.TimeFormatFunc, // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	hook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	core := zapcore.NewCore(encoder, hook, level)
	// infoHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path)))
	// errorHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path+".error")))
	//
	// infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl < zapcore.ErrorLevel })
	// errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= zapcore.ErrorLevel })
	//
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(encoder, infoHook, infoLevel),
	// 	zapcore.NewCore(encoder, errorHook, errorLevel),
	// )

	// 开启开发模式，堆栈跟踪
	// caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// skip := zap.AddCallerSkip(1)

	return zap.New(core, development).Sugar()
}
