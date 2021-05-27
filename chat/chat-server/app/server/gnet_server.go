package server

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet"
)

type NetServer struct {
	*gnet.EventServer
	handlerF func(data *[]byte, con gnet.Conn) error
	ip       string
	port     uint16
	network  string
	ctx      context.Context
	cancel   context.CancelFunc
}

func (n *NetServer) Setup(network, ip string, port uint16, f func(data *[]byte, con gnet.Conn) error) error {
	n.handlerF = f
	n.ip = ip
	n.port = port
	n.network = network
	n.ctx, n.cancel = context.WithCancel(context.Background())
	return nil
}

func (n *NetServer) Run() {
	err := gnet.Serve(n, fmt.Sprintf("%s//%s:%d", n.network, n.ip, n.port), gnet.WithMulticore(true), gnet.WithReusePort(true))
	if err != nil {
		panic(err)
	}
}

func (n *NetServer) Stop() {
	n.cancel()
}

func (n *NetServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func (n *NetServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (n *NetServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func newNetServer() *NetServer {
	server := new(NetServer)
	return server
}
