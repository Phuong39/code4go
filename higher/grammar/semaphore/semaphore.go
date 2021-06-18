package main

import (
	"fmt"
	"time"
)

func main() {
	var flag = true

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("exit!")
		flag = false
	}()

	for flag {
		fmt.Println("run!")
	}
}
