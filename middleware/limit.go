package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"
)

var bucket *ratelimit.Bucket

func InitRateLimit(rateLimitNum int) {
	bucket = ratelimit.NewBucket(time.Second/time.Duration(int64(rateLimitNum)), int64(rateLimitNum))
}

func RateLimit(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, req *http.Request) {
		if takeAvailable(bucket, 1) {
			next.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(f)
}

// takeAvailable 从 token bucket 中获取一个 token
func takeAvailable(bucket *ratelimit.Bucket, num int64) bool {

	// 配置为正数
	leftCount := bucket.TakeAvailable(num)

	return leftCount > 0
}
