package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"log"
	"sync"
	"time"
)

var dataMap = &sync.Map{}

func init() {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			count := 0
			dataMap.Range(func(key, value interface{}) bool {
				count++
				return true
			})
			fmt.Println("当前接受请求：", count)
		}
	}()
}

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)
	// Use ants pool to unblock the event-loop.
	_ = es.pool.Submit(func() {
		dataMap.Store(string(data), string(data))
		//if err := c.AsyncWrite(data); err != nil {
		//	panic(err)
		//}
	})
	return
}

func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &echoServer{pool: p}
	// WithMulticore(true) 代表设置多核处理
	log.Fatal(gnet.Serve(echo, "udp://:8080", gnet.WithMulticore(true)))
}
