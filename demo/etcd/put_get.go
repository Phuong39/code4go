package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
基本的 key-value 存取
*/
func put(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := client.Put(ctx, key, value)
	cancel()
	if err != nil {
		panic(err)
	}
}

/**
基本的 key-value 存取
*/
func get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		panic(err)
	}
	for _, ev := range resp.Kvs {
		return string(ev.Value), nil
	}
	return "", errors.New("not find!")
}

/**
定时的 key-value 存取
timeOut（秒）后自动执行DELETE事件
*/
func lease(key, value string, timeOut int64) {
	resp, err := client.Grant(context.TODO(), timeOut)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}
}

func keepAlive(key, value string) {
	resp, err := client.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}

	ch, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		panic(err)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
