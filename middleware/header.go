package middleware

// Warning: header middleware 必须放在其他 middleware 之前

import (
	"github.com/xuruiray/go-web-framework/util/request"
	"net/http"
	"strconv"
)

//HeaderSetter 请求头中间件
func HeaderSetter(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		record := &request.LogRecord{
			TraceID: strconv.Itoa(request.GenerateLogID()),
			Host:    req.Host,
			Method:  req.Method,
			Remote:  req.RemoteAddr,
		}

		ctx = request.SetLogRecord(ctx, record)
		ctx = request.SetRequest(ctx, req)

		next.ServeHTTP(w, req.WithContext(ctx))
	}
	return http.HandlerFunc(f)
}
