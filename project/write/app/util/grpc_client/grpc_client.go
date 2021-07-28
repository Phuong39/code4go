package grpc_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"utils"
)

func CreateServiceListenConn(ctx context.Context) *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9901", ctx)
}

func createGrpcClient(serviceAddress string, ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithUnaryInterceptor(utils.ClientInterceptorLowerReaches(utils.DefTracer, ctx)))
	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
