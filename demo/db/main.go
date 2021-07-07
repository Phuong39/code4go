package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.140.128:6379")
	if err != nil {
		panic(err)
	}

	fmt.Println(conn)
}
