package main

import (
	"fmt"
	"sync"
	"time"
)

/**
happens-before
假设A和B表示一个多线程的程序执行的两个操作。
如果A happens-before B，那么A操作对内存的影响将对执行B的线程(且执行B之前)可见。
*/

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	fmt.Println(count)
	wg.Wait()
}
