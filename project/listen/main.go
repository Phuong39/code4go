package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"listen/app/controller/listen_controller"
	"listen/app/proto/listen"
	"log"
	"net"
	"os"
	"utils"
)

const (
	ServiceName     = "gRPC-Service-Listen"
	ServiceHostPort = "0.0.0.0:9901"
)

func main() {
	var serviceOpts []grpc.ServerOption

	tracer, _, err := utils.NewJaegerTracer(ServiceName, utils.JaegerHostPort)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		serviceOpts = append(serviceOpts, utils.ServerTracerOption())
	}

	l, err := net.Listen("tcp", ServiceHostPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(serviceOpts...)

	// 服务注册
	listen.RegisterListenServer(s, &listen_controller.ListenController{})

	log.Println("Listen on " + ServiceHostPort)
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
