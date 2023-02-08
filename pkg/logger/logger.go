// Package logger 日志记录设置相关组件
package logger

import (
	"os"

	"gin_web/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Lg *zap.Logger

const (
	RELEASE = "release"
	TEST    = "test"
)

// NewLogger 初始化 logger
func NewLogger() (err error) {
	// 根据传入的日记级别进行记录
	var l = new(zapcore.Level)
	// 将配置文件的日志级别转化为zapcore.Level类型
	// DEBUG：详细的信息,通常只出现在诊断问题上
	// INFO：确认一切按预期运行
	// WARNING：一个迹象表明,一些意想不到的事情发生了,或表明一些问题在不久的将来(例如。磁盘空间低”)。这个软件还能按预期工作。
	// ERROR：更严重的问题,软件没能执行一些功能
	// CRITICAL：一个严重的错误,这表明程序本身可能无法继续运行
	if err = l.UnmarshalText([]byte(global.AppSetting.Log.Level)); err != nil {
		return
	}
	// 根据配置文件的是不是 release 来判断是否输出到终端
	var core zapcore.Core
	// writeSyncerError := getLogWriterError()
	// // 根据不同的 runmode 级别，来判断日志输出位置
	// switch global.ServerSetting.RunMode {
	// case "release":
	// 	core = zapcore.NewTee(
	// 		zapcore.NewCore(encoder, writeSyncer, l),
	// 		// 错误日志单独输出
	// 		zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
	// 	)
	// case "test":
	// 	core = zapcore.NewCore(encoder, writeSyncer, l)
	// default:
	// 	// 初始化终端输出配置
	// 	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	// 	// 输出一组，这样，终端和文件都输出
	// 	core = zapcore.NewTee(
	// 		zapcore.NewCore(encoder, writeSyncer, l),
	// 		// 错误日志单独输出
	// 		zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
	// 		// 输出到终端
	// 		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	// 	)
	// }
	// 根据 RunMode 选择输出位置
	switch global.ServerSetting.RunMode {
	// 生产和测试环境，输出到文件并且错误信息单独输出
	case RELEASE, TEST:
		encoder := getEncoder()
		writeSyncer := getLogWriter()
		writeSyncerError := getLogWriterError()
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			// 错误日志单独输出
			zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
		)
	default:
		// 初始化终端输出配置
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	}
	Lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Lg)
	return
}

// 格式化输出编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 正常输出，按照配置文件要求
func getLogWriter() zapcore.WriteSyncer {
	// fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + time.Now().Format("2006-01-02 15:04:05") + global.AppSetting.LogFileExt
	// 测试后发现不能添加时间，不然每次调用都新生成一个文件
	// accessFileName := global.AppSetting.LogSavePath + "/access_" + global.AppSetting.LogFileName + "_" + time.Now().Format("2006-01-02 15:04:05") + global.AppSetting.LogFileExt
	accessFileName := global.AppSetting.Log.LogSavePath + "/" + global.AppSetting.Log.LogFileName + "_" + global.AppSetting.Log.LogFileExt

	lumberJackLogger := &lumberjack.Logger{
		Filename:   accessFileName,                   // 日志文件的位置
		MaxSize:    global.AppSetting.Log.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: global.AppSetting.Log.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     global.AppSetting.Log.MaxAge,     // 保留旧文件的最大天数，不管副本数量有多少，超过这个时间的日志都删除
		LocalTime:  global.AppSetting.Log.LocalTime,  // LocalTime确定时间是否用于格式化中时间戳
		Compress:   global.AppSetting.Log.Compress,   // Compress决定是否对存储的日志文件进行压缩，使用gzip默认情况下不执行压缩。
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 错误日志单独输出
func getLogWriterError() zapcore.WriteSyncer {
	// fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + "Error" + time.Now().Format("2006-01-02 15:04:05") + global.AppSetting.LogFileExt
	errorFileName := global.AppSetting.Log.LogSavePath + "/error_" + global.AppSetting.Log.LogFileName + "_" + global.AppSetting.Log.LogFileExt
	lumberJackLogger := &lumberjack.Logger{
		Filename:   errorFileName,                    // 日志文件的位置
		MaxSize:    global.AppSetting.Log.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: global.AppSetting.Log.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     global.AppSetting.Log.MaxAge,     // 保留旧文件的最大天数，不管副本数量有多少，超过这个时间的日志都删除
		LocalTime:  global.AppSetting.Log.LocalTime,  // LocalTime确定时间是否用于格式化中时间戳
		Compress:   global.AppSetting.Log.Compress,   // Compress决定是否对存储的日志文件进行压缩，使用gzip默认情况下不执行压缩。
	}
	return zapcore.AddSync(lumberJackLogger)
}
