package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/xuruiray/go-web-framework/util/config"
	"time"
)

var redisClient *redis.Pool

func Init(file string) (err error) {

	var cacheConfig config.CacheConfig

	err = config.LoadConfig(file, &cacheConfig)
	if err != nil {
		return err
	}

	redisClient = &redis.Pool{
		MaxIdle:         cacheConfig.MaxIdle,
		MaxActive:       cacheConfig.MaxActive,
		IdleTimeout:     time.Duration(cacheConfig.IdleTimeoutMs) * time.Millisecond,
		MaxConnLifetime: time.Duration(cacheConfig.MaxConnLifeTimeMs) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cacheConfig.IP)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return err
}

func Get(ctx context.Context, key string) (string, error) {
	rc := redisClient.Get()
	defer rc.Close()

	return redis.String(rc.Do("get", key))
}

func Set(ctx context.Context, key string) (string, error) {
	rc := redisClient.Get()
	defer rc.Close()

	return redis.String(rc.Do("set", key))
}
