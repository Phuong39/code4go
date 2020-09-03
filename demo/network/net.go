package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		temp := make([]byte, 10*1024)
		conn.Read(temp)
		fmt.Println("-->", string(temp))

		conn.Write(temp)
	}
}
