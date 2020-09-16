package main

import (
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	client, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	go func() {
		var buff bytes.Buffer
		for {
			for {
				b := make([]byte, 1024*10)
				_, _ = client.Read(b)
				buff.WriteString(string(b))
			}
			fmt.Println(buff.String())
			buff.Reset()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		client.Write([]byte("hah"))
		time.Sleep(1 * time.Second)
		wg.Done()
	}
	wg.Wait()
}
