package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"jaegers/rpc/client"
	"log"
	"sync"
	"testing"
	"utils"
)

type HelloApiRpcClient struct {
	conn   *grpc.ClientConn
	Client client.HelloApiClient
}

var (
	defaultClient *HelloApiRpcClient
	once          = &sync.Once{}
)

func NewHelloApiRpcClient() *HelloApiRpcClient {
	once.Do(func() {
		//	建立RPC连接
		conn, err := grpc.Dial(RpcHost, grpc.WithInsecure(), grpc.WithUnaryInterceptor(utils.ClientInterceptor(utils.DefTracer, context.Background())))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defaultClient = &HelloApiRpcClient{conn: conn, Client: client.NewHelloApiClient(conn)}
	})
	return defaultClient
}

func TestNewMyApiRpcClient(t *testing.T) {
	_, closer, err := utils.NewJaegerTracer("rpc_client", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	rpc := NewHelloApiRpcClient()
	reply, err := rpc.Client.SayHello(context.Background(), &client.SayHelloReq{
		Name: "lsm",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
