package grammar

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Go 中定义的Happens Before保证

参考文档：https://studygolang.com/articles/14129

1.单线程
	在单线程环境下，所有的表达式，按照代码中的先后顺序，具有Happens Before关系
2.init函数
	如果包P1中导入了包P2，则P2中的init函数Happens Before 所有P1中的操作
	main函数Happens After 所有的init函数
3.Goroutine
	Goroutine的创建Happens Before所有此Goroutine中的操作
	Goroutine的销毁Happens After所有此Goroutine中的操作
4.Channel
	对一个元素的send操作Happens Before对应的receive 完成操作
	对channel的close操作Happens Before receive 端的收到关闭通知操作
	对于Unbuffered Channel，对一个元素的receive 操作Happens Before对应的send完成操作
	对于Buffered Channel，假设Channel 的buffer 大小为C，那么对第k个元素的receive操作，Happens Before第k+C个send完成操作。可以看出上一条Unbuffered Channel规则就是这条规则C=0时的特例
5.Lock
	对于一个Mutex/RWMutex，设n < m，则第n个Unlock操作Happens Before第m个Lock操作。
	对于一个RWMutex，存在数值n，RLock操作Happens After 第n个UnLock，其对应的RUnLockHappens Before 第n+1个Lock操作。
6.Once
	once.Do中执行的操作，Happens Before 任何一个once.Do调用的返回
*/

func TestSemaphore(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// 使用chan完成类似于信号量的功能，每次只有一个协程操作
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestVisible(t *testing.T) {
	var flag = true

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("exit!")
		flag = false
	}()

	for flag {
		fmt.Println("run!")
	}
}

func TestGoCreate(t *testing.T) {
	happensBeforeMulti(0)
}

// Sample Routine 1
func happensBeforeMulti(i int) {
	i += 2      // E1
	go func() { // G1 goroutine create
		fmt.Println(i) // E2
	}() // G2 goroutine destryo
}
