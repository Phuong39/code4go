package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	var cond = sync.NewCond(&lock)
	fmt.Println("hello world!!!")
	lock.Lock()
	go func() {
		cond.Wait()
		fmt.Println("123")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		cond.Signal()
	}()

	time.Sleep(10 * time.Second)
	lock.Unlock()
}
