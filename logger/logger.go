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

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/wuruipeng404/scaffold/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func timeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(util.TimeFormatString))
}

func newSugarLogger(path string) *zap.SugaredLogger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     timeFormat, // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	infoHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path)))
	errorHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path+".error")))

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl < zapcore.ErrorLevel })
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= zapcore.ErrorLevel })

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoHook, infoLevel),
		zapcore.NewCore(encoder, errorHook, errorLevel),
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	skip := zap.AddCallerSkip(1)

	return zap.New(core, caller, development, skip).Sugar()
}

// 根据日期分割日志
func logWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
