package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(conn, `GET / HTTP/1.0\r\n\r\n`)

	// status, err := bufio.NewReader(conn).ReadString('\n')
	// fmt.Println(status, err)

	b := make([]byte, 1)
	for {
		fmt.Println(conn.Read(b))
		fmt.Fprintf(conn, `G`)
	}
}
