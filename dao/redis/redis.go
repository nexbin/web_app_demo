package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 声明一个全局rdb变量
var (
	rdb *redis.Client
	ctx = context.Background()
)

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		DB:       viper.GetInt("redis.db"),
		Password: viper.GetString("redis.password"),
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
