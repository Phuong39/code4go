package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	count := 18000
	wg := &sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			if err := sendOnce([]byte(fmt.Sprintf("hello %d", i))); err != nil {
				panic(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func sendOnce(data []byte) error {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return err
	}
	defer socket.Close()

	// 发送数据
	_, err = socket.Write(data)
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return err
	}
	return nil
}
