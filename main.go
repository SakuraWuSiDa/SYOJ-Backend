package main

import (
	"context"
	"fmt"
	"github.com/XGHXT/SYOJ-Backend/config"
	"github.com/XGHXT/SYOJ-Backend/logger"
	"github.com/XGHXT/SYOJ-Backend/pkg"
	"github.com/XGHXT/SYOJ-Backend/pkg/mysql"
	"github.com/XGHXT/SYOJ-Backend/pkg/redis"
	"github.com/XGHXT/SYOJ-Backend/router"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 初始化日志
	if err := logger.Init(config.Config.LogConfig); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	// 初始化 mysql 连接
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}

	// 初始化 redis 连接
	if err := redis.Init(config.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}

	// 注册路由
	r := router.Setup(config.Config.Mode)

	// 启动服务 （优雅关机操作）
	server := &http.Server{
		Addr:    config.Config.Port,
		Handler: r,
	}

	// 初始化化 gin 框架内置的校验器使用的翻译器
	if err := pkg.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err: %v\n", err)
		return
	}

	// 开启一个 goroutine 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}