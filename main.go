package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/xuruiray/go-web-framework/rpc"
	"github.com/xuruiray/go-web-framework/storage/mysql"
	"github.com/xuruiray/go-web-framework/storage/redis"
	"github.com/xuruiray/go-web-framework/util/config"
	"github.com/xuruiray/go-web-framework/util/server"
)

func main() {
	var (
		err error
	)

	logo()

	go http.ListenAndServe(":8015", nil)

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

	fmt.Println("init rpc")
	err = rpc.Init(config.RPCConfigFile)
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

func logo() {
	println(`	
	 _   _      _ _        __        __         _     _
	| | | | ___| | | ___   \ \      / /__  _ __| | __| |
	| |_| |/ _ \ | |/ _ \   \ \ /\ / / _ \| |__| |/ _| |
	|  _  |  __/ | | (_) |   \ V  V / (_) | |  | | (_| |
	|_| |_|\___|_|_|\___/     \_/\_/ \___/|_|  |_|\____|`)
}
