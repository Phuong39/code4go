package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
	_ "yidu-user/config"
	"yidu-user/handler"
	proto "yidu-user/protos"
)

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Version("latest"),
		micro.Name("yidu-user"),
		micro.Registry(etcd.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"127.0.0.1:2380"}
		})),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// 解析命令行等初始化
	service.Init()

	// 注册 handler
	proto.RegisterUserHandler(service.Server(), new(handler.UserHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
