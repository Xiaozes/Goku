package Glogin

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var DB *redis.Client

func RedisLogin(host, port, password string) bool {
	rdb := redis.NewClient(&redis.Options{
		// 需要修改成你的配置，本地无需修改
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接成功")
	// 成功连接将其赋值给全局变量
	DB = rdb
	return true
}
