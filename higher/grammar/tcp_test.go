package grammar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"runtime"
	"testing"
)

func TestManyAccept(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:8090")
	assert.NoError(t, err)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				conn, err := listen.Accept()
				assert.NoError(t, err)
				bytes := make([]byte, 0, 1024)
				le, err := conn.Read(bytes)
				assert.NoError(t, err)
				fmt.Println(bytes[:le])
			}
		}()
	}
	// 网络层接收完数据包后递交给UDP后，UDP的处理过程。该过程UDP需要做的工作就是接收数据包并对其进行校验，校验成功后将其放入接收队列中等待用户空间程序来读取。
	// 用户空间程序调用read()等系统调用读取已经放入接收队列中的数据。
}
