package rpc

import (
	"context"
	"testing"
	"utils"
)

func TestStartRPCServer(t *testing.T) {
	_, closer, err := utils.NewJaegerTracer("rpc_server", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	StartRPCServer(context.Background())
}
