package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"read/app/controller/read_controller"
	"read/app/proto/read"
	"utils"
)

const (
	ServiceName     = "gRPC-Service-Read"
	ServiceHostPort = "0.0.0.0:9903"
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
	read.RegisterReadServer(s, &read_controller.ReadController{})

	log.Println("Listen on " + ServiceHostPort)
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
