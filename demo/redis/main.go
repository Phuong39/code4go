package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient *redis.Pool
)

func init() {
	RedisClient = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "119.29.117.244:6379")
			if err != nil {
				return nil, err
			}
			c.Do("auth", "fuck123")
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}

func main() {
	conn := RedisClient.Get()
	defer conn.Close()
	_, _ = conn.Do("set", "key", "hello")
	val, err := conn.Do("get", "key")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", val)
}
