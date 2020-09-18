/**
* 作者：刘时明
* 时间：2019/10/20-15:58
* 作用：
 */
package conf

import (
	"fmt"
	"novel/utils"
	"strings"
)

const (
	// SvcName 注册到consul上的服务名
	SvcName = "net.novel.api.demo"
)

type RedisConfig struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type MySqlConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type AppConfig struct {
	IsDev bool
	Redis RedisConfig `json:"redis"`
	Mysql MySqlConfig `json:"mysql"`
}

var C AppConfig

func InitConfig() {
	array := strings.Split(SvcName, ".")
	if len(array) == 0 {
		panic("InitConfig error , no service name")
	}
	sName := array[len(array)-1]
	err := utils.InitConf(sName, &C)
	if err != nil {
		panic(fmt.Sprintf("InitConfig failed %v", err))
	}
	fmt.Println("配置信息=", C)
}
