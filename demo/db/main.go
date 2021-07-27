package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var list []int

	list = nil

	fmt.Println(len(list))

	list = append(list, 1)

	fmt.Println(list)

	list2 := make([]int, 0, 10)
	fmt.Printf("%p\n", list2)

	list2 = append(list2, 1)

	demo(list2)
	fmt.Printf("result %p\n", list2)
	fmt.Println(list2)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	exchangeSlice(slice)
	fmt.Println(slice)
}

func demo(l []int) {
	l = append(l, 10, 100)
	fmt.Printf("demo %p\n", l)
}

func exchangeSlice(slice []int) {
	slice = slice[0:0]
}
