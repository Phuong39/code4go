package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

const HOST = "127.0.0.1:2379"

var client *clientv3.Client

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to etcd success")
	client = cli
}

func main() {
	key := "key_1"
	go watch(key, func(event *clientv3.Event) {
		fmt.Println("event => ", event)
	})

	// lease(key, "hello", 2)
	put(key, "hello")
	// keepAlive(key, "hello")

	fmt.Println(get(key))

	lock()
}
