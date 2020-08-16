/*
* 作者：刘时明
* 时间：2020/8/16 0016-12:36
* 作用：
 */
package concurrent

import (
	"fmt"
	"runtime"
)

func SetAndGetProc() {
	fmt.Println("当前线程数量=", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(10)
	fmt.Println("设置后=", runtime.GOMAXPROCS(0))
}
