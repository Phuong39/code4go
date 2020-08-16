/*
* 作者：刘时明
* 时间：2020/8/16 0016-23:27
* 作用：
 */
package memory

import (
	"fmt"
	"os"
)
import "github.com/shirou/gopsutil/process"

var ps *process.Process

func GetMemoryInfo(n int) {
	if ps == nil {
		p, err := process.NewProcess(int32(os.Getpid()))
		if err != nil {
			panic(err)
		}
		ps = p
	}
	info, _ := ps.MemoryInfo()
	fmt.Printf("%d. VMS:%d MB,RSS:%d MB \n", n, info.VMS>>20, info.RSS>>20)
}

func ShowMemoryInfo() {
	GetMemoryInfo(1)

	data := new([10][1024]byte)

	GetMemoryInfo(2)

	for i := range data {
		for x, n := 0, len(data[i]); x < n; x++ {
			data[i][x] = 1
		}
		GetMemoryInfo(3)
	}
}
