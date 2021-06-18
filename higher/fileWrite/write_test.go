package fileWrite

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

const Count = 1000 * 10000

const LogVal = "abcdefghiejkimlopqrstuvwxyzabcdefghiejkimlopqrstuvwxyzabcdefghiejkimlopqrstuvwxyz\n"

var filename = "demo.log"

var wg = sync.WaitGroup{}

func TestWriteFile(t *testing.T) {
	t1 := time.Now()
	var err error
	var file *os.File
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	writer := bufio.NewWriter(file)

	for i := 0; i < Count; i++ {
		_, err = writer.Write([]byte(strconv.Itoa(i) + LogVal))
		assert.NoError(t, err)
		if i%1000 == 0 {
			_ = writer.Flush()
		}
	}
	fmt.Println(time.Now().Unix() - t1.Unix())
}

func TestWriteQueue(t *testing.T) {
	t1 := time.Now()
	queue := NewWriteQueue(1024)
	go queue.Consumer()
	wg.Add(Count)
	for i := 0; i < Count/1000; i++ {
		go func(i int) {
			for j := 0; j < 1000; j++ {
				queue.Append(strconv.Itoa(i*1000+j) + LogVal)
				wg.Done()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Unix() - t1.Unix())
}

func TestWriteCASQueue(t *testing.T) {

}

type WriteQueue struct {
	val chan string
}

func NewWriteQueue(cap int) *WriteQueue {
	return &WriteQueue{
		val: make(chan string, cap),
	}
}

func (w *WriteQueue) Append(val string) {
	w.val <- val
}

func (w *WriteQueue) Consumer() {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for {
		val, ok := <-w.val
		if !ok {
			break
		}
		_, _ = writer.WriteString(val)
	}
}
