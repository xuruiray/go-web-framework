package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn

func Init() {
	var err error
	conn, err = redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer conn.Close()
}

func Get(ctx context.Context, key string) (string, error) {
	return redis.String(conn.Do("get", key))
}

func Set(ctx context.Context, key string) (string, error) {
	return redis.String(conn.Do("set", key))
}
