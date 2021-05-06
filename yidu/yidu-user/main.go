package main

import (
	"github.com/micro/go-micro/v2"
	"yidu-user/handler"
	proto "yidu-user/protos"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("yidu-user"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterUserHandler(service.Server(), new(handler.UserHandler))

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
