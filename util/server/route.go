package server

import (
	"github.com/go-chi/chi"
	"github.com/xuruiray/go-web-framework/controller"
	"github.com/xuruiray/go-web-framework/middleware"
)

func initRouter() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middleware.RateLimit)
	mux.Use(middleware.HeaderSetter)

	// 查询特征依赖关系
	mux.Handle("/hello", &BaseHandler{controller.Hello})

	return mux
}
