package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
	"os"
	"syscall"
	"unsafe"
)

func mmapTest() {
	at, err := mmap.Open("/Users/xyz/test_mmap_data.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	buff := make([]byte, 10)
	//读入的长度为slice预设的长度，0是offset。预设长度过长将会用0填充。
	at.ReadAt(buff, 0)
	fmt.Println(string(buff))
	at.Close()
}

func mmapSocket() {
	n := 1000
	t := int(unsafe.Sizeof(0)) * n
	map_file, _ := os.Create("/tmp/test.dat")
	_, _ = map_file.Seek(int64(t-1), 0)
	map_file.Write([]byte(" "))
	mmap, _ := syscall.Mmap(map_file.Fd(), 0, t, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	map_array := (*[1000]int)(unsafe.Pointer(&mmap))
	for i := 0; i < n; i++ {
		map_array[i] = i * i
	}

	fmt.Println(*map_array)
}
