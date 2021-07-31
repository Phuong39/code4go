package main

import (
	"context"
	"github.com/PegasusMeteor/grpc-examples/common"
	helloworld "github.com/PegasusMeteor/grpc-examples/proto/consul"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	port        = ":50051"
	jaegerAgent = "192.168.140.128:6831"
	serviceName = "HelloServer"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	tracer, closer, err := common.NewJaegerTracer(serviceName, jaegerAgent)
	defer closer.Close()
	if err != nil {
		log.Printf("NewJaegerTracer err: %v", err.Error())
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(common.ServerInterceptor(tracer)))
	helloworld.RegisterGopherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
