package server

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"syscall"

	"github.com/go-chi/chi"

	"github.com/xuruiray/binding"
	"github.com/xuruiray/go-web-framework/middleware"
	"github.com/xuruiray/go-web-framework/util/config"
	"github.com/xuruiray/go-web-framework/util/convert"
	"github.com/xuruiray/go-web-framework/util/logger"
)

// Init : 初始化 web server
func Init(port string) error {

	// 初始化 限流
	middleware.InitRateLimit(config.MainConfig.RateLimit)

	// 初始化 route
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
				logger.Warnf(context.TODO(), "shutdown http server failed")
			}
		}
	}()

	return nil
}

type BaseHandler struct {
	handleFunc interface{}
}

// Request req
type Request interface{}

// Response resp
type Response interface{}

func (bh *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request, err := bh.bindRequest(r)
	if err != nil {
		logger.Errorf(r.Context(), "bind request failed|%v", err)
	}

	returnValues := reflect.ValueOf(bh.handleFunc).
		Call([]reflect.Value{reflect.ValueOf(r.Context()), reflect.ValueOf(request)})

	response := returnValues[0].Interface()

	if err := responseJSON(r.Context(), w, response); err != nil {
		logger.Errorf(r.Context(), "response json failed|%v", err)
	}
}

// responseJSON 处理控制层返回的结构体 为 json字符串，不处理返回结构体的data包装
func responseJSON(ctx context.Context, w http.ResponseWriter, respData Response) error {

	jsonStr, err := json.Marshal(respData)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	// 将返回数据发送到客户端
	_, err = w.Write(jsonStr)
	return err
}

func (bh *BaseHandler) bindRequest(r *http.Request) (request Request, err error) {
	request = reflect.New(reflect.TypeOf(bh.handleFunc).In(1).Elem()).Interface()
	if err = bind(r, request); err == nil {
		return
	}
	return
}

// Bind 将 http 的 request body，或 form 绑定到 dst 的 interface
func bind(r *http.Request, dst interface{}) (err error) {
	contentType := contentType(r)
	// 默认为 form
	if len(contentType) == 0 {
		contentType = "application/x-www-form-urlencoded"
	}

	err = binding.Binder(r.Method, contentType).Bind(r, dst)
	err = bindRoute(r, dst)

	return
}

func bindRoute(req *http.Request, dst interface{}) error {
	// 转换 restful 参数至 http 参数中
	urlParams := chi.RouteContext(req.Context()).URLParams
	for i, key := range urlParams.Keys {
		req.Form[key] = append(req.Form[key], urlParams.Values[i])
	}

	return convert.MapToStructByTag(req.Form, dst, "route")
}

func contentType(r *http.Request) string {
	var ct string
	if values, ok := r.Header["Content-Type"]; ok {
		ct = values[0]
	}

	for i, char := range ct {
		if char == ' ' || char == ';' {
			ct = ct[:i]
		}
	}
	return ct
}
