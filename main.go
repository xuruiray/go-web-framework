package main

import (
	"fmt"
	"github.com/xuruiray/go-web-framework/storage/mysql"
	"github.com/xuruiray/go-web-framework/storage/redis"
	"github.com/xuruiray/go-web-framework/util/config"
	"github.com/xuruiray/go-web-framework/util/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		err error
	)

	err = config.Init(config.AppConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = mysql.Init(config.DataBaseConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = redis.Init(config.CacheConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = server.Init(config.MainConfig.Port)
	if err != nil {
		fmt.Println(err)
	}

	initSignal()
	// wait for system quit
	<-systemQuit
}

var systemQuit = make(chan int, 1)

//initSignal 初始化信号
func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGALRM, syscall.SIGTERM, syscall.SIGUSR1)

	// Block until a signal is received.
	<-c
	systemQuit <- 1
}
