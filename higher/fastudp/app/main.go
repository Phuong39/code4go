package main

import (
	"fmt"
	"github.com/shaoyuan1943/fastudp"
	"net"
	"sync/atomic"
	"time"
)

var count uint32

type MyEventHandler struct {
}

func (e *MyEventHandler) OnReaded([]byte, *net.UDPAddr) {
	atomic.AddUint32(&count, 1)
}

func main() {
	go countSpeed()
	server, err := fastudp.NewUDPServer("udp", "192.168.74.1:8500", 10, new(MyEventHandler))
	if err != nil {
		panic(err)
	}
	for !server.IsClosed() {
	}
}

func countSpeed() {
	for {
		tmp := count
		time.Sleep(5 * time.Second)
		fmt.Println("speed in 5 seconds is ", (count-tmp)/5.0)
	}
}
