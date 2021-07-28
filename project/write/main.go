package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"utils"
	"write/app/controller/write_controller"
	"write/app/proto/write"
)

const (
	ServiceName     = "gRPC-Service-Write"
	ServiceHostPort = "0.0.0.0:9904"
)

func main() {
	var serviceOpts []grpc.ServerOption

	tracer, closer, err := utils.NewJaegerTracer(ServiceName, utils.JaegerHostPort)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		serviceOpts = append(serviceOpts, utils.ServerTracerOption())
	}
	defer closer.Close()

	l, err := net.Listen("tcp", ServiceHostPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(serviceOpts...)

	// 服务注册
	write.RegisterWriteServer(s, &write_controller.WriteController{})

	log.Println("Listen on " + ServiceHostPort)
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
