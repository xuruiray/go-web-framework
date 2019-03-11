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

	fmt.Println("init app")
	err = config.Init(config.AppConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("init mysql")
	err = mysql.Init(config.DataBaseConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("init redis")
	err = redis.Init(config.CacheConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("init http server")
	err = server.Init(config.MainConfig.Port)
	if err != nil {
		fmt.Println(err)
	}

	// wait for system quit
	initSignal()
}

//initSignal 初始化信号
func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGALRM, syscall.SIGTERM, syscall.SIGUSR1)
	sign := <-c
	fmt.Println("app terminated by signal", sign)
}
