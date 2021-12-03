package main

// @title JS流量统计管理系统后台API
// @version 1.0
// @description JS流量统计管理系统后台API

// @contact.name xiayoushuang
// @contact.email york-xia@gmail.com

// @schemes http https
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

import (
	"js_statistics/app"
	"js_statistics/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	// _ "github.com/mkevac/debugcharts"
)

func main() {
	cfg := config.GetConfig()
	// go monitor.Start()
	go app.Run(cfg.Server.Port)
	go app.RunJs(cfg.JsServer.Port)

	// 性能监控
	// go http.ListenAndServe(":7090", nil)

	s := waiForSignal()
	log.Fatalf("signal received, server closed, %v", s)
}

func waiForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
