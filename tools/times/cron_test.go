package times

import (
	"fmt"
	"testing"
)

func TestNewTimingJob(t *testing.T) {
	// 每秒打印一个hello
	NewTimingJob("*/1 * * * * ?", func() {
		fmt.Println("hello")
	})
	select {}
}
