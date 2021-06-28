package fileWrite

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"syscall"
	"testing"
	"unsafe"
)

func TestZeroCopy(t *testing.T) {
	err := MappingRead(func(bytes []byte) {
		fmt.Println(string(bytes))
	})
	assert.NoError(t, err)

	err = UDPServer(8888, func(bytes []byte) {
		fmt.Println(string(bytes))
	})
	assert.NoError(t, err)
}

func MappingRead(readF func([]byte)) error {
	fd, err := syscall.Open("a.txt", syscall.GENERIC_ALL, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)
	fSize, err := syscall.Seek(fd, 0, 2)
	syscall.Seek(fd, 0, 0)
	// 每个block的页面个数
	var pageLock int64 = 8192
	// 数据块数
	var bLockSize = int64(syscall.Getpagesize()) * pageLock
	var bLockNum = fSize / bLockSize
	if fSize%bLockSize > 0 {
		bLockNum = bLockNum + 1
	}
	hMap, err := syscall.CreateFileMapping(fd, nil, syscall.PAGE_READONLY, uint32(fSize>>32), uint32(fSize), nil)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(hMap)
	var i int64
	for i = 0; i < bLockNum; i++ {
		var dataLen = int32(bLockSize)
		if i == (bLockNum - 1) {
			dataLen = int32(fSize % bLockSize)
		}
		// 开始读取内存块
		var currPos = i * bLockSize
		var tmpLen = int(dataLen)
		addr, err := syscall.MapViewOfFile(hMap, syscall.FILE_MAP_READ, 0, uint32(currPos), uintptr(tmpLen))
		if err != nil {
			return err
		}
		// 把addr变成slice
		d := (*[1 << 28]byte)(unsafe.Pointer(addr))
		readF(d[:tmpLen])
		err = syscall.UnmapViewOfFile(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func UDPServer(port uint16, handler func([]byte)) error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return err
	}
	sa := &syscall.SockaddrInet4{
		Port: int(port),
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Bind(fd, sa)
	if err != nil {
		return err
	}
	err = syscall.Listen(fd, 0)
	if err != nil {
		return err
	}

	for {
		var buf []byte
		le, err := syscall.Read(fd, buf)
		if err != nil {
			return err
		}
		handler(buf[:le])
	}
}
