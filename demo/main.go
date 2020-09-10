/*
* 作者：刘时明
* 时间：2020/8/16 0016-11:59
* 作用：
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var arr = []chan int{make(chan int, 1), make(chan int, 1)}

func main() {
	//channel.CloseExit()
	//
	//concurrent.SetAndGetProc()
	//
	//concurrent.LocalStorage()
	//
	//concurrent.GoExit()

	// memory.ShowMemoryInfo()
	go B()
	go C()
	go A()

	time.Sleep(10 * time.Second)
}

func A() {
	<-arr[0]
	fmt.Println("a")
}

func B() {
	select {
	case <-arr[0]:
	case <-arr[1]:
	}
	fmt.Println("b")
	arr[0] <- 10
}

func C() {
	fmt.Println("c")
	arr[1] <- 10
}
