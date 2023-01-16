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
	"github.com/aloysZy/gin_web/pkg/setting"
)

func init() {
	// 初始化配置文件
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
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

func main() {
	fmt.Printf("global.ServerSetting:%#v\nglobal.AppSeting:%#v\nglobal.DatabaseSetting.Mysql:%#v\n", global.ServerSetting, global.AppSetting, global.DatabaseSetting.Mysql)
	router := routers.NewRouter()
	// 不使用 run 启动，自定义配置服务参数
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 开启一个goroutine启动服务
	go func() {
		log.Printf("[%v]ListenAndServe: %v Actual pid is %d", "gin_web", s.Addr, syscall.Getpid())
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
