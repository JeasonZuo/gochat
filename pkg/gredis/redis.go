package gredis

import (
	"context"
	"github.com/JeasonZuo/gochat/global"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func SetUp() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("Redis.Host"),
		Password: viper.GetString("Redis.Password"), // no password set
		DB:       viper.GetInt("Redis.DefaultDB"),   // use default DB
		PoolSize: viper.GetInt("Redis.PoolSize"),
	})

	result := rdb.Ping(ctx)
	if result.Val() != "PONG" {
		panic("redis connect error")
	}

	global.Redis = rdb
}
