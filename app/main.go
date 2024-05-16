package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Start a TCP server that listens on port 4221
	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("Listening on port 4221")

	// Accept a connection and handle it
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
		return
	}
	defer conn.Close()

	// Read from the connection
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	req := string(buff[:n])
	lines := strings.Split(req, "\r\n")
	reqLine := strings.Split(lines[0], " ")
	path := reqLine[1]

	// Write to the connection
	if path == "/" {
		responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("Hello World"), "Hello World")
		conn.Write([]byte(responce))
	}
}
