package study

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const count = 5

var ch chan *string

var wg sync.WaitGroup

func init() {
	ch = make(chan *string, 5)
}

func TestChanCloseRead(t *testing.T) {
	wg.Add(2)
	go func() {
		for i := 0; i < count; i++ {
			time.Sleep(1 * time.Second)
			str := fmt.Sprintf("%s%d", "hello", i)
			ch <- &str
			fmt.Println("写入一次。。。")
		}
		close(ch)
		fmt.Println("写入退出")
		wg.Done()
	}()

	go func() {
		for i := 0; i < count+1; i++ {
			str := <-ch
			fmt.Println("读取一次, str=", str)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main exit.")
}
