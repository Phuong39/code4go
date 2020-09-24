package collection

import (
	"fmt"
	"testing"
)

type Int int

func (i Int) Compare(o interface{}) int {
	t, ok := o.(Int)
	if !ok {
		return 0
	}
	return int(i) - int(t)
}

func TestQueue(t *testing.T) {
	queue := new(PriorityDeque)
	queue.Add(Int(1))
	queue.Add(Int(2))
	queue.Add(Int(-10))
	fmt.Println(queue)
}
