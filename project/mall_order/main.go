package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"mall_order/client"
	"mall_order/service"
	"net"
	"utils"
)

func main() {
	_, closer, err := utils.NewJaegerTracer("mall_user", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	var defaultServer = &service.OrderApiServerImpl{}
	rpcServer := grpc.NewServer(utils.ServerTracerOption())
	// 允许客户端获取server中已注册的服务和方法等信息
	client.RegisterOrderApiServer(rpcServer, defaultServer)
	reflection.Register(rpcServer)
	defer rpcServer.GracefulStop()
	lis, err := net.Listen("tcp", utils.OrderHost)
	if err != nil {
		panic(err)
	}
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
