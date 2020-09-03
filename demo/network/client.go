package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			temp := make([]byte, 1024*10)
			le, err := conn.Read(temp)
			if err != nil {
				panic(err)
			}
			if le > 0 {
				fmt.Println("收到回复=>", string(temp[0:le]))
			}
		}
	}()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		conn.Write([]byte("hello"))
		time.Sleep(50 * time.Millisecond)
		fmt.Println("发送一次")
		wg.Done()
	}
	wg.Wait()
}
