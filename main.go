package main

// @title JS流量统计管理系统后台API
// @version 1.0
// @description JS流量统计管理系统后台API

// @contact.name xiayoushuang
// @contact.email york-xia@gmail.com

// @schemes http https
// @BasePath /

import (
	"js_statistics/app"
	"js_statistics/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	go app.Run(cfg.Server.Port)
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
