package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"time"
)

const LOCK_KEY = "lock_key"

/**
实现分布式锁
*/
func lock() {
	client1, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer client1.Close()
	client2, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer client2.Close()

	s1, err := concurrency.NewSession(client1)
	if err != nil {
		panic(err)
	}
	m1 := concurrency.NewMutex(s1, LOCK_KEY)

	s2, err := concurrency.NewSession(client2)
	if err != nil {
		panic(err)
	}
	m2 := concurrency.NewMutex(s2, LOCK_KEY)

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("released lock for s1")
	<-m2Locked
	fmt.Println("acquired lock for s2")
}
