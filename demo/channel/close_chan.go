/*
* 作者：刘时明
* 时间：2020/8/16 0016-12:22
* 作用：
 */
package channel

import "time"

/**
关闭通道解除阻塞
 */
func CloseExit() {
	exit := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")
		close(exit)
	}()
	println("main...")
	<-exit
	println("main exit...")
}
