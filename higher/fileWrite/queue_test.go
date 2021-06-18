package fileWrite

import (
	"fmt"
	"testing"
)

func TestCASQueue(t *testing.T) {
	queue := NewCASQueue(1024)
	_ = queue.Put("str")
	val, _ := queue.Get()
	fmt.Println(val)
}
