package server

import (
	"context"
	"github.com/xuruiray/go-web-framework/util/logger"
	"net/http"
	"os"
	"syscall"
)

// Init : 初始化 web server
func Init(port string) error {

	mux := initRouter()

	if port[0] != ':' {
		port = ":" + port
	}

	pid := os.Getpid()
	go func() {
		if err := http.ListenAndServe(port, mux); err != nil {
			process := os.Process{Pid: pid}
			err := process.Signal(syscall.SIGINT)
			if err != nil {
				logger.Warn(context.TODO(), "shutdown http server failed")
			}
		}
	}()

	return nil
}
