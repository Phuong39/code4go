package base

import (
	"fmt"
	"testing"
)

/**
Slice 截取操作不会拷贝底层数组，而是共享
*/
func TestSliceSub(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}

	arr2 := arr1[0:1]
	// arr1和arr2地址不同，但是它们底层的数组指针是相同的
	fmt.Printf("%p - %p \n", &arr1, &arr2)

	arr2 = append(arr2, 100)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
