package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strings"
)

var logger *zap.Logger

func Init(filePath string, fs ...zapcore.Field) {
	hook := lumberjack.Logger{
		Filename:   filePath, // 日志文件路径
		MaxSize:    128,      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,       // 日志文件最多保存多少个备份
		MaxAge:     7,        // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(fs...)
	// 构造日志
	logger = zap.New(core, caller, development, filed)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, getCaller(&fields)...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, getCaller(&fields)...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, getCaller(&fields)...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, getCaller(&fields)...)
}

func getCaller(fields *[]zap.Field) []zap.Field {
	_, file, line, _ := runtime.Caller(2)
	arr := strings.Split(file, "/")
	l := len(arr)
	*fields = append(*fields, zap.String("call", fmt.Sprintf("%v/%v:%v", arr[l-2], arr[l-1], line)))
	return *fields
}
