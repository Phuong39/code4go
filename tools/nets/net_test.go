package nets

import (
	"fmt"
	"testing"
	"time"
)

func TestScanPort(t *testing.T) {
	udpPorts := ScanPort("udp", "127.0.0.1", 0, 65535, time.Second*3)
	tcpPorts := ScanPort("tcp", "127.0.0.1", 0, 65535, time.Second*3)
	fmt.Println("开放的UDP端口：", udpPorts)
	fmt.Println("开放的TCP端口：", tcpPorts)
}
