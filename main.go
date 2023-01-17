package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/internal/routers"
	"github.com/aloysZy/gin_web/pkg/logger"
	"github.com/aloysZy/gin_web/pkg/setting"
)

func init() {
	// 初始化配置文件
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	// 初始化日志
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
	// 初始化 MySQL
	if err := setupMysqlDBEngin(); err != nil {
		log.Fatalf("init.setupMysqlDBEngin err:%v", err)
	}
}

// setupSetting初始化配置文件
func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 输入的这个key，就是配置文件中的 key
	if err := newSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err := newSetting.ReadSection("APP", &global.AppSetting); err != nil {
		return err
	}
	if err := newSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	// 默认是纳秒，将传入的时间转化为秒
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// 初始化日志
func setupLogger() error {
	if err := logger.NewLogger(); err != nil {
		return err
	}
	return nil
}

// setupMysqlDBEngin 初始化 MySQL
func setupMysqlDBEngin() error {
	// 这里一定要是"=",初始化全局变量,不然其他包调用的时候依然是 nil
	var err error
	global.MysqlDBEngine, err = model.NewMysqlDBEngine(global.DatabaseSetting.Mysql)
	if err != nil {
		return err
	}
	return nil
}

// @title gin_web
// @version 1.0
// @description 练习 Gin 写 web 服务
// @termsOfService https://github.com/aloysZY/gin_web
func main() {
	fmt.Printf("global.ServerSetting:%#v\nglobal.AppSeting:%#v\nglobal.DatabaseSetting.Mysql:%#v\n", global.ServerSetting, global.AppSetting, global.DatabaseSetting.Mysql)
	router := routers.NewRouter()
	// 不使用 run 启动，自定义配置服务参数
	// https://blog.csdn.net/yanyuan_smartisan/article/details/113357813
	// ReadTimeout覆盖了从连接被接受到请求body被完全读取的时间（如果你确实读取了正文，否则就是到header结束）。它的实现是在net/http中，通过在Accept后立即调用SetReadDeadline来实现的。s
	// WriteTimeout通过在readRequest结束时调用SetWriteDeadline来设置，通常覆盖从请求header读取结束到响应写入结束（也就是ServeHTTP的生存期）的时间。
	// WriteTimeout通过在readRequest结束时调用SetWriteDeadline来设置，通常覆盖从请求header读取结束到响应写入结束（也就是ServeHTTP的生存期）的时间。[参考]
	// 但是，当连接是HTTPS时，SetWriteDeadline在Accept之后立即被调用，以便它也覆盖作为TLS握手的一部分而写入的数据包。令人烦恼的是，这意味着（只在这种情况下）WriteTimeout最终包括读取头部和等待第一个字节的时间。[参考]
	// 在处理不受信任的客户端和/或网络时，应该将这两个超时都设置上，以便客户端无法通过慢读写来长时间持有连接。
	// 最后是http.TimeoutHandler。它不是服务器参数，而是一个限制ServeHTTP调用最大duration的Handler包装器。它的工作方式是缓冲响应，并在超过最后期限时发送504 Gateway Timeout。注意，它在1.6中被破坏并在1.6.2中得到修复。[参考]
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 开启一个goroutine启动服务
	go func() {
		log.Printf("[%v]ListenAndServe%v Actual pid is %d", global.ServerSetting.Name, s.Addr, syscall.Getpid())
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shuting down server...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
