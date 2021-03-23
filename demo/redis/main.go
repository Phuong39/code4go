package main

import (
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
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			// c.Do("auth", "fuck123")
			// c.Do("SELECT", 0)
			return c, nil
		},
	}
}

func main() {
	conn := RedisClient.Get()
	time.Sleep(10 * time.Second)
	_, _ = conn.Do("set", "key", "hello")
	//val, err := conn.Do("get", "key")
	//if err != nil {
	//	panic(err)
	//}
	defer conn.Close()

	// fmt.Printf("%s", val)
}
