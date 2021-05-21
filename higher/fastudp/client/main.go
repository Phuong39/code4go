package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var count uint32

func countSpeed() {
	for {
		tmp := count
		time.Sleep(5 * time.Second)
		fmt.Println("speed in 5 seconds is ", (count-tmp)/5.0)
	}
}

func main() {
	go countSpeed()
	for {
		sendJob()
	}
}

var sendData = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func sendJob() {
	addr := net.UDPAddr{
		IP:   []byte{192, 168, 74, 1},
		Port: 8500,
	}
	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100000; i++ {
		sendOnce(conn)
	}
}

func sendOnce(conn *net.UDPConn) {
	_, err := conn.Write(sendData)
	if err != nil {
		panic(err)
	}
	atomic.AddUint32(&count, 1)
}
