/*
* 作者：刘时明
* 时间: 2019/9/30-10:42
* 作用：配置信息
 */
package utils

import (
	"github.com/micro/go-micro/config"
	"io"
	"os"
)

const (
	DefaultConfigPath = "/conf/conf.json"
	ConsulRegister    = iota
	EtcdRegister
	MDNSRegister
	TCPTransport = iota
	HTTPTransport
	GRPCTransport
	NotBroker = iota
	HTTPBroker
	RabbitBroker
	NsqBroker
	NatsBroker
	RedisBroker
)

// Option 可选参数设置函数
type Option func(o *Options)

// Options 配置的可选参数
type Options struct {
	AppName        string // 应用名称
	ServerName     string // 服务名称
	EncoderType    string // 编码方式
	WatchAppCommon bool   // 是否监听App层下的公用路径
	LogOutput      io.Writer
	RegisterType   int      // 注册中心类型
	RegisterAddr   []string // 注册中心地址
	ConfigPath     string   // 配置文件路径
	Transport      int      // 数据传输方式
	BrokerType     int      // 发布订阅推送类型
	MQAddr         []string // MQ地址
}

// 默认的配置
var defaultOptions = Options{
	EncoderType:  "json",            // 默认使用json编码
	LogOutput:    os.Stdout,         // 默认日志输出到标准输出
	RegisterType: ConsulRegister,    // 默认使用Consul
	ConfigPath:   DefaultConfigPath, // 配置文件默认在conf目录的conf.json
	Transport:    GRPCTransport,     // 默认使用grpc
	BrokerType:   NotBroker,         // 默认使用HTTP推送
	MQAddr:       []string{},        // MQ地址默认为空
}

// InitConf 初始化
func InitConf(serverName string, conf interface{}, opts ...Option) error {
	defaultOptions.AppName = AppName
	defaultOptions.ServerName = serverName
	defaultOptions.ConfigPath = "./" + serverName + defaultOptions.ConfigPath
	for _, v := range opts {
		v(&defaultOptions)
	}
	if err := loadConfig(conf, &defaultOptions); err != nil {
		return err
	}
	return nil
}

// loadConfig 加载配置文件
func loadConfig(conf interface{}, opt *Options) error {
	if err := config.LoadFile(opt.ConfigPath); err != nil {
		return err
	}
	return config.Scan(conf)
}
