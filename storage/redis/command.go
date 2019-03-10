package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/xuruiray/go-web-framework/util/config"
)

var conn redis.Conn

func Init(file string) error {
	var (
		err         error
		cacheConfig config.CacheConfig
	)

	err = config.LoadConfig(file, &cacheConfig)
	if err != nil {
		return err
	}

	conn, err = redis.Dial("tcp", ":6379")
	return err
}

func Get(ctx context.Context, key string) (string, error) {
	return redis.String(conn.Do("get", key))
}

func Set(ctx context.Context, key string) (string, error) {
	return redis.String(conn.Do("set", key))
}
