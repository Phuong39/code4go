package main

import (
	"context"
	"github.com/PegasusMeteor/grpc-examples/common"
	helloworld "github.com/PegasusMeteor/grpc-examples/proto/consul"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	jaegerAgent   = "192.168.140.128:6831"
	serviceName   = "HelloClient"
)

func main() {
	tracer, closer, err := common.NewJaegerTracer(serviceName, jaegerAgent)
	defer closer.Close()
	if err != nil {
		panic(err)
	}
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.1:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(common.ClientInterceptor(tracer)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGopherClient(conn)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
		cancel()
		time.Sleep(5 * time.Second)
	}
}
