package collection

import (
	"unsafe"
)

/**
* 作者：刘时明
* 时间：2019/9/11-11:35
* 作用：
 */

type HashMap struct {
	len       int
	entryList []LinkedList
}

type Entry struct {
	key   interface{}
	value interface{}
}

func (h *HashMap) Put(key, value interface{}) interface{} {
	hash := hashCode(key, cap(h.entryList))
	list := h.entryList[hash]
	return list
}

func hashCode(o interface{}, cap int) int {
	hash := uintptr(unsafe.Pointer(&o))
	return int(hash) % cap
}
