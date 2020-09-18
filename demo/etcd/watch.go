package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

/**
使用Watch可以监听键值对
*/
func watch(key string, handle func(event *clientv3.Event)) {
	// 返回一个 WatchResponse channel
	rch := client.Watch(context.Background(), key)
	for temp := range rch {
		for _, ev := range temp.Events {
			handle(ev)
		}
	}
}
