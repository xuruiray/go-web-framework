package server

import (
	"github.com/go-chi/chi"
	"github.com/xuruiray/go-web-framework/controller"
)

func initRouter() *chi.Mux {
	mux := chi.NewMux()

	// 查询特征依赖关系
	mux.Handle("/hello", &BaseHandler{controller.Hello})

	return mux
}
