package rpc

import (
	"context"
	"fmt"
	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"jaegers/rpc/client"
	"net"
	"os"
	"utils"
)

const RpcHost = ":9008"

type HelloApiServerImpl struct {
}

func (h *HelloApiServerImpl) SayHello(ctx context.Context, req *client.SayHelloReq) (*client.SayHelloReply, error) {
	reply := &client.SayHelloReply{}
	reply.Code = 200
	reply.Data = "hello: " + req.Name
	return reply, nil
}

func StartRPCServer(ctx context.Context) {
	var defaultServer = &HelloApiServerImpl{}
	rpcServer := grpc.NewServer(utils.ServerTracerOption())
	// 允许客户端获取server中已注册的服务和方法等信息
	client.RegisterHelloApiServer(rpcServer, defaultServer)
	reflection.Register(rpcServer)
	go func() {
		<-ctx.Done()
		fmt.Println("[GRPCServer]Received context done signal")
		rpcServer.GracefulStop()
		fmt.Println("[GRPCServer]Exiting")
	}()
	lis, err := net.Listen("tcp", RpcHost)
	if err != nil {
		log.Fatalf("failed listen. err:%v", err)
		os.Exit(1)
	}
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed Serve. err:%v", err)
		os.Exit(1)
	}
}
