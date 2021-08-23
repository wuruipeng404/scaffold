/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/8/20 17:21
 */

package logger

import (
	"io"
	"os"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "gopkg.in/natefinch/lumberjack.v2"
)

const (
	TimeFormatString = "2006-01-02 15:04:05.000"
)

func timeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimeFormatString))
}

func NewSugarLogger(path string) *zap.SugaredLogger {

	// hook := lumberjack.Logger{
	// 	Filename:   path,  // 日志文件路径
	// 	MaxSize:    300,   // 每个日志文件保存的最大尺寸 单位：M
	// 	MaxBackups: 10,    // 日志文件最多保存多少个备份
	// 	MaxAge:     7,     // 文件最多保存多少天
	// 	Compress:   false, // 是否压缩
	// 	LocalTime:  true,
	// }

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

	// 日志输出格式
	// if cLog.Encoding == "json" {
	// encoder := zapcore.NewJSONEncoder(encoderConfig)
	// } else {
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// }

	// 日志输出 文件,控制台
	// if config.Config.Debug {
	// infoHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	infoHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path)))
	errorHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter(path+".error")))
	// } else {
	// infoHook := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	// }

	infoLevel := zap.LevelEnablerFunc(func(le zapcore.Level) bool {
		return le < zapcore.ErrorLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(le zapcore.Level) bool {
		return le >= zapcore.ErrorLevel
	})

	// logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoHook, infoLevel),
		zapcore.NewCore(encoder, errorHook, errorLevel),
	)

	// zap.Config{}
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段

	// filed := zap.Fields(zap.Any("_Node", ""))
	// 构造日志

	// sugar log 性能稍微差一些,但更符合使用习惯
	return zap.New(core, development, caller).Sugar()
}

// 根据日期分割日志
func logWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d", // No go style anti-human format
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
