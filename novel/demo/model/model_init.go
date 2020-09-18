/**
* 作者：刘时明
* 时间：2019/10/21-22:32
* 作用：
 */
package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"novel/demo/conf"
	"xorm.io/xorm"
)

var (
	redisClient *redis.Client
	mysqlEngine *xorm.Engine
)

func Init() {
	redisClient = newRedisClient()
	mysqlEngine = newMysqlEngine()
}

func newRedisClient() (RedisClient *redis.Client) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.C.Redis.Addr,
		PoolSize: 20,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic("RedisInit error!")
	}
	fmt.Println("pong=>", pong)
	return
}

func newMysqlEngine() (mysqlEngine *xorm.Engine) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.C.Mysql.UserName, conf.C.Mysql.Password, conf.C.Mysql.Addr, conf.C.Mysql.Port, conf.C.Mysql.DBName)
	mysqlEngine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		panic(err)
	}
	return
}
