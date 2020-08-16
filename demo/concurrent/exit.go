/*
* 作者：刘时明
* 时间：2020/8/16 0016-12:55
* 作用：
 */
package concurrent

import (
	"fmt"
	"runtime"
	"time"
)

func GoExit() {
	go func() {
		defer func() {
			fmt.Println("step 4...")
		}()
		fmt.Println("step 1...")
		fmt.Println("step 2...")
		// 终止当前任务，但会确保已注册的延迟调用被执行
		runtime.Goexit()
		fmt.Println("step 3...")
	}()
	time.Sleep(2 * time.Second)
}
