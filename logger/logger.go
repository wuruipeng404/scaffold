/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/20 17:21
 */

package logger

import (
	"io"
	"log"
	"os"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/wuruipeng404/scaffold/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TimeFormatFunc(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(util.TimeFormatMill))
}

func NewSugarLogger(option *InitOption) *zap.SugaredLogger {

	if option == nil {
		option = &InitOption{
			loglevel: zap.DebugLevel,
			writers:  []io.Writer{os.Stdout},
			skip:     1,
		}
	} else {
		if len(option.writers) == 0 {
			option.writers = append(option.writers, os.Stdout)
		}
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
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
	})

	var syncer []zapcore.WriteSyncer

	for _, i := range option.writers {
		syncer = append(syncer, zapcore.AddSync(i))
	}

	hook := zapcore.NewMultiWriteSyncer(syncer...)
	level := zap.NewAtomicLevelAt(option.loglevel)

	core := zapcore.NewCore(encoder, hook, level)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	dev := zap.Development()
	skip := zap.AddCallerSkip(option.skip)

	return zap.New(core, caller, dev, skip).Sugar()
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
		log.Fatalf("rotate logger error:%s", err)
	}
	return hook
}

func NewRotateLogger(filename string) io.Writer {
	return logWriter(filename)
}
