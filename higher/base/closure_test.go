package base

import (
	"fmt"
	"sync"
	"testing"
)

/**
闭包
*/

func TestClosure(t *testing.T) {
	a := 1
	b := 2
	c := executeFn(func() int {
		a = a + b
		return a
	})
	fmt.Println(a, ",", b, ",", c)

	goFunc()
}

func executeFn(fn func() int) int {
	return fn()
}

func goFunc() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
}
