/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/20 17:21
 */

package logger

import (
	"io"
	"os"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/wuruipeng404/scaffold/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TimeFormatFunc(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(util.TimeFormatString))
}

func newSugarLogger(path string, skip int) *zap.SugaredLogger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     TimeFormatFunc, // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	hook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path)))
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
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	skipOpt := zap.AddCallerSkip(skip)

	return zap.New(core, caller, development, skipOpt).Sugar()
}

// 根据日期分割日志
func logWriter(filename string) io.Writer {
	hook, err := rotateLogs.New(
		filename+".%Y-%m-%d",
		rotateLogs.WithLinkName(filename),
		rotateLogs.WithMaxAge(time.Hour*24*7),
		rotateLogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
