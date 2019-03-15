package request

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

// ContentKey 打日志的 key
const ContentKey = "logRecordContextKey"
const RequestKey = "requestContextKey"

type keyType string

// LogRecord 请求信息
type LogRecord struct {
	URL       string
	Host      string
	Method    string
	TraceID   string
	Remote    string
	SrcMethod string
}

// LogID : request id
var LogID int64

// 初始化LogID
func init() {
	LogID = time.Now().UnixNano()
}

// GenerateLogID 获取logId
func GenerateLogID() int64 {
	return atomic.AddInt64(&LogID, 1)
}

//Set 将 k-v 存入 context
func Set(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, keyType(key), value)
}

//Get 从 Context 获取 key 对应的 value
func Get(ctx context.Context, key string) (interface{}, error) {
	result := ctx.Value(keyType(key))
	if result == nil {
		return 0, fmt.Errorf("can`t find %s in context", key)
	}
	return result, nil
}

// SetLogRecord 将 k-v 存入 context
func SetLogRecord(ctx context.Context, record *LogRecord) context.Context {
	return context.WithValue(ctx, ContentKey, record)
}

// GetLogRecord 获取LogRecord信息
func GetLogRecord(ctx context.Context) (record *LogRecord) {
	result := ctx.Value(ContentKey)
	if record, ok := result.(*LogRecord); ok {
		return record
	}

	return &LogRecord{}
}

// SetRequest 将 req 存入 context
func SetRequest(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, RequestKey, req)
}

// GetRequest 将 req 取出
func GetRequest(ctx context.Context, req *http.Request) *http.Request {
	result := ctx.Value(RequestKey)
	if req, ok := result.(*http.Request); ok {
		return req
	}
	return nil
}

// String logRecord to string
func (r *LogRecord) String() string {
	return fmt.Sprintf("traceid=%s|url=%s|method=%s|host=%s|host=%s|remote=%s",
		r.TraceID, r.URL, r.Method, r.Host, r.Remote)
}
