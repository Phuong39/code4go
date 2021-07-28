package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"mall_pay/client"
	"mall_pay/service"
	"net"
	"utils"
)

func main() {
	_, closer, err := utils.NewJaegerTracer("mall_pay", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	var defaultServer = &service.PayApiServerImpl{}
	rpcServer := grpc.NewServer(utils.ServerTracerOption())
	// 允许客户端获取server中已注册的服务和方法等信息
	client.RegisterPayApiServer(rpcServer, defaultServer)
	reflection.Register(rpcServer)
	defer rpcServer.GracefulStop()
	lis, err := net.Listen("tcp", utils.PayHost)
	if err != nil {
		panic(err)
	}
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
