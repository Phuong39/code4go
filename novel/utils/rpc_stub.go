/*
* 作者：刘时明
* 时间: 2019/9/30-10:41
* 作用：
 */
package utils

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/micro/go-micro/service/grpc"
	tsGrpc "github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-micro/transport/http"
	"github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/broker/redis"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/transport/tcp"
	"runtime"
	"time"
)

var innerService micro.Service

// 初始化服务
func InitService(svcName string, opts ...Option) (micro.Service, error) {
	for _, v := range opts {
		v(&defaultOptions)
	}
	var innerOpts []micro.Option
	innerOpts = append(innerOpts, micro.Name(svcName),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10))
	// 注册中心
	innerOpts = append(innerOpts, initRegister())
	// 数据传输方式
	innerOpts = append(innerOpts, initTransport())
	// 发布订阅推送方式
	temp := initBroker()
	if temp != nil {
		innerOpts = append(innerOpts, temp)
	}
	innerService = grpc.NewService(innerOpts...)
	innerService.Init()
	// windows环境不需要初始化日志文件目录
	if runtime.GOOS != "windows" {
		InitLogPath(svcName)
	}
	return innerService, nil
}

// Start 启动框架
func Start() error {
	// Run service
	if err := innerService.Run(); err != nil {
		return err
	}
	return nil
}

// Client 返回底层client
func Client() client.Client {
	if innerService != nil {
		return innerService.Client()
	}
	return nil
}

// Stop 进程退出时调用
func Stop() {
	log.Close()
}

// 初始化注册中心
func initRegister() micro.Option {
	var register micro.Option
	switch defaultOptions.RegisterType {
	case ConsulRegister:
		if len(defaultOptions.RegisterAddr) != 0 {
			register = micro.Registry(consul.NewRegistry(func(options *registry.Options) {
				options.Addrs = defaultOptions.RegisterAddr
			}))
		} else {
			// 默认本机agent
			register = micro.Registry(consul.NewRegistry())
		}
	case EtcdRegister:
		if len(defaultOptions.RegisterAddr) != 0 {
			register = micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
				options.Addrs = defaultOptions.RegisterAddr
			}))
		} else {
			// 默认本机agent
			register = micro.Registry(etcdv3.NewRegistry())
		}
	case MDNSRegister:
		register = micro.Registry(mdns.NewRegistry())
	default:
		panic("不能识别的注册中心")
	}
	return register
}

// 初始化数据传输方式
func initTransport() micro.Option {
	var transport micro.Option
	switch defaultOptions.Transport {
	case TCPTransport:
		transport = micro.Transport(tcp.NewTransport())
	case HTTPTransport:
		transport = micro.Transport(http.NewTransport())
	case GRPCTransport:
		transport = micro.Transport(tsGrpc.NewTransport())
	default:
		panic("不能识别的传输方式")
	}
	return transport
}

// 初始化推送
func initBroker() micro.Option {
	var brokerTemp micro.Option
	var addrFun = func(options *broker.Options) {
		if len(defaultOptions.MQAddr) > 0 {
			options.Addrs = defaultOptions.MQAddr
		}
	}
	switch defaultOptions.BrokerType {
	case NotBroker:
		// 不使用发布订阅
		return nil
	case HTTPBroker:
		// 默认使用http，无需处理
		brokerTemp = nil
	case RabbitBroker:
		brokerTemp = micro.Broker(rabbitmq.NewBroker(addrFun))
	case NsqBroker:
		brokerTemp = micro.Broker(nsq.NewBroker(addrFun))
	case NatsBroker:
		brokerTemp = micro.Broker(nats.NewBroker(addrFun))
	case RedisBroker:
		brokerTemp = micro.Broker(redis.NewBroker(addrFun))
	default:
		panic("不能识别的推送方式")
	}
	err := broker.Init()
	if err != nil {
		panic(err)
	}
	err = broker.Connect()
	if err != nil {
		panic(err)
	}
	return brokerTemp
}
