package base

import (
	"fmt"
	"testing"
)

/**
for range会为每一个元素创建副本，每个副本共享一块内存
*/
func TestRange(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	// v的地址相同，也就是采用覆盖的方式
	for index, v := range arr {
		fmt.Printf("%p %p\n", &v, &arr[index])
	}
}
