package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // Redis服务器密码（如果有的话）
		DB:       0,                // 使用的数据库编号
	})

	ctx := context.Background()
	// 尝试与Redis建立连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	// 执行Redis命令
	err = client.Set(ctx, "key", "Hello", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)

	// 关闭Redis连接
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
