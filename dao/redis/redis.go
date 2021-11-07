package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"web_app/settings"
)

// 声明一个全局rdb变量
var (
	rdb *redis.Client
	ctx = context.Background()
)

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port,
		),
		DB:       config.DB,
		Password: config.Password,
		PoolSize: config.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
