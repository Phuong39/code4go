package nets

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/**
TCP 端口扫描程序
*/

func isOpen(network, host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout(network, fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func ScanPort(network, hostname string, startPort, endPort int, timeout time.Duration) []int {
	var ports []int
	wg := &sync.WaitGroup{}
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(network, hostname, p, timeout)
			if opened {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	return ports
}
