package base

import (
	"fmt"
	"testing"
)

/**
Slice本质是一个结构体，拥有len、cap和一个指向数组的指针
*/
func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	changeAndAppend(arr)
	fmt.Println(arr)

	clear(arr)
	fmt.Println(arr)
}

func changeAndAppend(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
	arr = append(arr, 10)
}

func clear(arr []int) {
	arr = arr[0:0]
}
