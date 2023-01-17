package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/aloysZy/gin_web/global"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

const (
	RELEASE = "release"
	TEST    = "test"
)

// NewLogger 初始化 logger
func NewLogger() (err error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var l = new(zapcore.Level)
	// 将配置文件的日志级别转化为zapcore.Level类型
	// DEBUG：详细的信息,通常只出现在诊断问题上
	// INFO：确认一切按预期运行
	// WARNING：一个迹象表明,一些意想不到的事情发生了,或表明一些问题在不久的将来(例如。磁盘空间低”)。这个软件还能按预期工作。
	// ERROR：更严重的问题,软件没能执行一些功能
	// CRITICAL：一个严重的错误,这表明程序本身可能无法继续运行
	if err = l.UnmarshalText([]byte(global.AppSetting.Level)); err != nil {
		return
	}
	// 根据配置文件的是不是 release 来判断是否输出到终端
	var core zapcore.Core
	// 根据不同的 runmode 级别，来判断日志输出位置
	/*	switch global.ServerSetting.RunMode {
		case "release":
			core = zapcore.NewTee(
				zapcore.NewCore(encoder, writeSyncer, l),
				// 错误日志单独输出
				zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
			)
		case "test":
			core = zapcore.NewCore(encoder, writeSyncer, l)
		default:
			// 初始化终端输出配置
			consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
			// 输出一组，这样，终端和文件都输出
			core = zapcore.NewTee(
				zapcore.NewCore(encoder, writeSyncer, l),
				// 错误日志单独输出
				zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
				// 输出到终端
				zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			)
		}*/
	switch global.ServerSetting.RunMode {
	case RELEASE:
		writeSyncerError := getLogWriterError()
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			// 错误日志单独输出
			zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel),
		)
	case TEST:
		core = zapcore.NewCore(encoder, writeSyncer, l)
	default:
		// 初始化终端输出配置
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	}
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

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
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + time.Now().Format("2006-01-02 15:04:05") + global.AppSetting.LogFileExt
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                     // 日志文件的位置
		MaxSize:    global.AppSetting.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: global.AppSetting.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     global.AppSetting.MaxAge,     // 保留旧文件的最大天数，不管副本数量有多少，超过这个时间的日志都删除
		LocalTime:  global.AppSetting.LocalTime,  // LocalTime确定时间是否用于格式化中时间戳
		Compress:   global.AppSetting.Compress,   // Compress决定是否对存储的日志文件进行压缩，使用gzip默认情况下不执行压缩。
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 错误日志单独输出
func getLogWriterError() zapcore.WriteSyncer {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + "Error" + time.Now().Format("2006-01-02 15:04:05") + global.AppSetting.LogFileExt
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                     // 日志文件的位置
		MaxSize:    global.AppSetting.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: global.AppSetting.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     global.AppSetting.MaxAge,     // 保留旧文件的最大天数，不管副本数量有多少，超过这个时间的日志都删除
		LocalTime:  global.AppSetting.LocalTime,  // LocalTime确定时间是否用于格式化中时间戳
		Compress:   global.AppSetting.Compress,   // Compress决定是否对存储的日志文件进行压缩，使用gzip默认情况下不执行压缩。
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志 stack 布尔值来记录对战信息
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
