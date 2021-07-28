package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"mall_user/client"
	"mall_user/service"
	"net"
	"utils"
)

func main() {
	_, closer, err := utils.NewJaegerTracer("mall_user", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	var defaultServer = &service.UserApiServerImpl{}
	rpcServer := grpc.NewServer(utils.ServerTracerOption())
	// 允许客户端获取server中已注册的服务和方法等信息
	client.RegisterUserApiServer(rpcServer, defaultServer)
	reflection.Register(rpcServer)
	defer rpcServer.GracefulStop()
	lis, err := net.Listen("tcp", utils.UserHost)
	if err != nil {
		panic(err)
	}
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
