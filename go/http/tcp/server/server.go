package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

var b = make([]byte, 1)

func handleConnection(conn net.Conn) {
	// 读取数据，处理
	for {
		fmt.Println(conn.Read(b))
		fmt.Println(string(b))
		conn.Write([]byte("H"))
	}
	// 短连接
	conn.Close()
}
