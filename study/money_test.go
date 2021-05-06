package study

import (
	"fmt"
	"github.com/bnkamalesh/currency"
	"testing"
)

func TestMoney(t *testing.T) {
	c1, _ := currency.New(50, 50, "INR", "10.50", "paise", 100)
	fmt.Println(c1.Float64())
	c1.AddInt(10, 0)
	fmt.Println(c1.Float64())
}
