package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("Listening on port 4221")

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
		return
	}
	defer conn.Close()

	contents := "Hello"
	responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(contents), contents)
	conn.Write([]byte(responce))
}
