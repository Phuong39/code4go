package study

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var val int64

func TestCAS(t *testing.T) {
	result := atomic.CompareAndSwapInt64(&val, 0, 100)
	fmt.Println(result)
	fmt.Println(val)
}
