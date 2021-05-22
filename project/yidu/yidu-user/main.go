package main

import (
	proto "commons/protos/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
	"yidu-user/config"
	_ "yidu-user/config"
	"yidu-user/handler"
)

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Version(config.C.Server.Version),
		micro.Name(config.C.Server.Name),
		micro.Registry(etcd.NewRegistry(func(options *registry.Options) {
			options.Addrs = config.C.Registry.Addrs
		})),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// 解析命令行等初始化
	service.Init()

	// 注册 handler
	err := proto.RegisterUserHandler(service.Server(), new(handler.UserHandler))
	if err != nil {
		panic(err)
	}

	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
