/*
* 作者：刘时明
* 时间：2020/8/16 0016-12:49
* 作用：
 */
package concurrent

import (
	"fmt"
	"sync"
)

func LocalStorage() {
	var wg sync.WaitGroup
	var gs [5]struct {
		id     int
		result int
	}
	for i := 0; i < len(gs); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}
	wg.Wait()
	fmt.Printf("%+v \n", gs)
}
