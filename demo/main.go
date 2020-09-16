/*
* 作者：刘时明
* 时间：2020/8/16 0016-11:59
* 作用：
 */
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Student struct {
	Name string
}

func main() {
	s := &Student{}
	of := reflect.TypeOf(s)
	fmt.Println(of)
	fmt.Println(of.Name())
	fmt.Println(of.Kind())
	fmt.Println(of.Elem())

	e := of.Elem()
	fmt.Println(e.Size())
	fmt.Println(of.Size())
	fmt.Println(unsafe.Sizeof(*s))
}
